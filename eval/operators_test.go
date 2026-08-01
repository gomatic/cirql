package eval

import (
	"errors"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

func TestEval_Comparisons(t *testing.T) {
	cases := []struct {
		l    value.Value
		r    value.Value
		op   ast.BinOp
		want bool
	}{
		{op: ast.OpGt, l: int64(3), r: int64(2), want: true},
		{op: ast.OpLt, l: int64(1), r: int64(2), want: true},
		{op: ast.OpGe, l: int64(2), r: int64(2), want: true},
		{op: ast.OpLe, l: int64(2), r: int64(2), want: true},
		{op: ast.OpGt, l: int64(1), r: int64(2), want: false},
	}
	for _, c := range cases {
		got := evalOK(t, ast.BinaryExpr{Op: c.op, L: lit(c.l), R: lit(c.r)}, Env{})
		if got != c.want {
			t.Fatalf("%v %s %v = %v want %v", c.l, c.op, c.r, got, c.want)
		}
	}
}

func TestEval_Compare_TypeError(t *testing.T) {
	_, err := Eval(ast.BinaryExpr{Op: ast.OpGt, L: lit("a"), R: lit(int64(1))}, Env{})
	if !errors.Is(err, value.ErrIncomparable) {
		t.Fatalf("got %v want ErrIncomparable", err)
	}
}

func TestEval_Arithmetic(t *testing.T) {
	// int op int: +, -, * preserve int64 (one consistent model across the
	// additive and multiplicative operators); / and % use the double model.
	cases := []struct {
		want value.Value
		op   ast.BinOp
	}{
		{op: ast.OpAdd, want: int64(5)},
		{op: ast.OpSub, want: int64(1)},
		{op: ast.OpMul, want: int64(6)},
		{op: ast.OpDiv, want: 1.5},
		{op: ast.OpMod, want: 1.0},
	}
	for _, c := range cases {
		got := evalOK(t, ast.BinaryExpr{Op: c.op, L: lit(int64(3)), R: lit(int64(2))}, Env{})
		if got != c.want {
			t.Fatalf("3 %s 2 = %v (%T) want %v (%T)", c.op, got, got, c.want, c.want)
		}
	}
}

// A mixed int/float operand promotes to float for every arithmetic operator —
// no operator silently keeps an int when a float is involved.
func TestEval_ArithmeticMixedPromotesToFloat(t *testing.T) {
	for _, op := range []ast.BinOp{ast.OpAdd, ast.OpSub, ast.OpMul} {
		got := evalOK(t, ast.BinaryExpr{Op: op, L: lit(int64(3)), R: lit(2.0)}, Env{})
		if _, ok := got.(float64); !ok {
			t.Errorf("3 %s 2.0 = %v (%T), want a float64", op, got, got)
		}
	}
}

// Integer subtraction and multiplication stay exact past 2^53, where the float
// model would round — the reason to preserve int64 rather than promote.
func TestEval_IntArithmeticIsExactPast2Pow53(t *testing.T) {
	big := int64(9007199254740993) // 2^53 + 1, unrepresentable as float64
	got := evalOK(t, ast.BinaryExpr{Op: ast.OpMul, L: lit(big), R: lit(int64(1))}, Env{})
	if got != big {
		t.Errorf("got %v, want exact %d", got, big)
	}
}

func TestEval_Arith_TypeErrors(t *testing.T) {
	if _, err := Eval(ast.BinaryExpr{Op: ast.OpSub, L: lit("x"), R: lit(int64(1))}, Env{}); !errors.Is(err, ErrType) {
		t.Fatalf("left type: got %v want ErrType", err)
	}
	if _, err := Eval(ast.BinaryExpr{Op: ast.OpSub, L: lit(int64(1)), R: lit("x")}, Env{}); !errors.Is(err, ErrType) {
		t.Fatalf("right type: got %v want ErrType", err)
	}
}

func TestEval_Logical_ShortCircuit(t *testing.T) {
	boom := ast.FuncCall{Name: "nope"}
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpAnd, L: lit(false), R: boom}, Env{}); got != false {
		t.Fatalf("false&&boom = %v want false", got)
	}
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpOr, L: lit(true), R: boom}, Env{}); got != true {
		t.Fatalf("true||boom = %v want true", got)
	}
}

func TestEval_Logical_Full(t *testing.T) {
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpAnd, L: lit(true), R: lit(false)}, Env{}); got != false {
		t.Fatalf("true&&false = %v want false", got)
	}
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpOr, L: lit(false), R: lit(true)}, Env{}); got != true {
		t.Fatalf("false||true = %v want true", got)
	}
}

func TestEval_Logical_Errors(t *testing.T) {
	boom := ast.FuncCall{Name: "nope"}
	if _, err := Eval(ast.BinaryExpr{Op: ast.OpAnd, L: boom, R: lit(true)}, Env{}); !errors.Is(err, ErrUnknownFunc) {
		t.Fatalf("left err: got %v", err)
	}
	if _, err := Eval(ast.BinaryExpr{Op: ast.OpAnd, L: lit(true), R: boom}, Env{}); !errors.Is(err, ErrUnknownFunc) {
		t.Fatalf("right err: got %v", err)
	}
}

// TestCmpResultRefusesAnOperatorThatMakesNoOrderingClaim covers the boundary the
// narrow type introduced. Before, the default arm answered EVERY unrecognised
// operator as "<=", so an operator added to the AST without a case here would
// have produced a plausible-looking boolean instead of an obvious failure —
// the worst kind of wrong answer, because a filter would silently keep the
// wrong objects.
func TestCmpResultRefusesAnOperatorThatMakesNoOrderingClaim(t *testing.T) {
	for _, c := range []ordering{-1, 0, 1} {
		if cmpResult(comparisonOp("undeclared"), c) {
			t.Fatalf("an undeclared operator claimed an ordering held at c=%d", c)
		}
	}

	if !cmpResult(cmpLessEqual, 0) {
		t.Fatal("<= must still hold at equality — the declared operators are unaffected")
	}
	if cmpResult(cmpLessEqual, 1) {
		t.Fatal("<= must not hold when the left side is greater")
	}
}

// TestArithFloatRefusesAnUndeclaredOperator covers the same boundary for
// arithmetic, where the old default returned the MODULUS of the operands for
// anything it did not recognise.
func TestArithmeticOpOutsideTheDeclaredSetYieldsNull(t *testing.T) {
	if got := arithFloat(arithmeticOp("undeclared"), 7, 2); got != nil {
		t.Fatalf("arithFloat(undeclared) = %v, want null rather than a silent modulus", got)
	}
	if got := arithFloat(arithMod, 7, 2); got != 1.0 {
		t.Fatalf("arithFloat(arithMod) = %v, want 1 — the declared operators still work", got)
	}
}
