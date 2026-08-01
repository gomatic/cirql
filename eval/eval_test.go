package eval

import (
	"errors"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

func lit(v value.Value) ast.Expr { return ast.Literal{V: v} }

func field(names ...string) ast.FieldAccess {
	fa := ast.FieldAccess{}
	for _, n := range names {
		fa.Path = append(fa.Path, ast.PathSegment{Name: n})
	}
	return fa
}

func objEnv(m map[string]value.Value) Env {
	return Env{Obj: m, Vars: map[string]value.Value{}}
}

func evalOK(t *testing.T, e ast.Expr, env Env) value.Value {
	t.Helper()
	v, err := Eval(e, env)
	if err != nil {
		t.Fatalf("Eval error: %v", err)
	}
	return v
}

func TestError_Message(t *testing.T) {
	if ErrType.Error() != "eval: type error" {
		t.Fatalf("unexpected message %q", ErrType.Error())
	}
}

func TestEval_Literal(t *testing.T) {
	if got := evalOK(t, lit(int64(5)), Env{}); got != int64(5) {
		t.Fatalf("got %v want 5", got)
	}
}

func TestEval_FieldAccess(t *testing.T) {
	env := objEnv(map[string]value.Value{"a": map[string]value.Value{"b": int64(7)}})
	if got := evalOK(t, field("a", "b"), env); got != int64(7) {
		t.Fatalf("got %v want 7", got)
	}
}

func TestEval_FieldAccess_Identity(t *testing.T) {
	env := objEnv(map[string]value.Value{"a": int64(1)})
	got := evalOK(t, ast.FieldAccess{}, env)
	if _, ok := got.(map[string]value.Value); !ok {
		t.Fatalf("identity = %T want object", got)
	}
}

func TestEval_FieldAccess_NullPropagation(t *testing.T) {
	env := objEnv(map[string]value.Value{})
	if got := evalOK(t, field("missing", "x"), env); got != nil {
		t.Fatalf("got %v want nil", got)
	}
}

func TestEval_FieldAccess_OnNonObject(t *testing.T) {
	env := Env{Obj: int64(3), Vars: map[string]value.Value{}}
	if got := evalOK(t, field("a"), env); got != nil {
		t.Fatalf("got %v want nil", got)
	}
}

func TestEval_FieldAccess_Iter(t *testing.T) {
	list := []value.Value{
		map[string]value.Value{"n": int64(1)},
		map[string]value.Value{"n": int64(2)},
	}
	env := objEnv(map[string]value.Value{"items": list})
	got := evalOK(t, ast.FieldAccess{Path: []ast.PathSegment{
		{Name: "items"}, {IsIter: true}, {Name: "n"},
	}}, env)
	out, ok := got.([]value.Value)
	if !ok || len(out) != 2 || out[0] != int64(1) || out[1] != int64(2) {
		t.Fatalf("iter map = %#v", got)
	}
}

func TestEval_FieldAccess_IterNonList(t *testing.T) {
	env := objEnv(map[string]value.Value{"items": int64(3)})
	got := evalOK(t, ast.FieldAccess{Path: []ast.PathSegment{
		{Name: "items"}, {IsIter: true},
	}}, env)
	if got != nil {
		t.Fatalf("iter non-list = %v want nil", got)
	}
}

func TestEval_VarRef(t *testing.T) {
	env := Env{Vars: map[string]value.Value{"x": int64(9)}}
	if got := evalOK(t, ast.VarRef{Name: "x"}, env); got != int64(9) {
		t.Fatalf("got %v want 9", got)
	}
	if got := evalOK(t, ast.VarRef{Name: "missing"}, env); got != nil {
		t.Fatalf("missing var = %v want nil", got)
	}
}

func TestEval_UnaryNot(t *testing.T) {
	if got := evalOK(t, ast.UnaryExpr{Op: ast.OpNot, X: lit(false)}, Env{}); got != true {
		t.Fatalf("!false = %v want true", got)
	}
}

func TestEval_UnaryNeg(t *testing.T) {
	if got := evalOK(t, ast.UnaryExpr{Op: ast.OpNeg, X: lit(int64(5))}, Env{}); got != int64(-5) {
		t.Fatalf("-5 = %v want -5", got)
	}
	if got := evalOK(t, ast.UnaryExpr{Op: ast.OpNeg, X: lit(2.5)}, Env{}); got != -2.5 {
		t.Fatalf("-2.5 = %v want -2.5", got)
	}
}

func TestEval_UnaryNeg_TypeError(t *testing.T) {
	_, err := Eval(ast.UnaryExpr{Op: ast.OpNeg, X: lit("x")}, Env{})
	if !errors.Is(err, ErrType) {
		t.Fatalf("got %v want ErrType", err)
	}
}

func TestEval_Unary_PropagatesError(t *testing.T) {
	_, err := Eval(ast.UnaryExpr{Op: ast.OpNot, X: ast.FuncCall{Name: "nope"}}, Env{})
	if !errors.Is(err, ErrUnknownFunc) {
		t.Fatalf("got %v want ErrUnknownFunc", err)
	}
}

func TestEval_EqNe(t *testing.T) {
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpEq, L: lit(int64(1)), R: lit(1.0)}, Env{}); got != true {
		t.Fatalf("1==1.0 = %v want true", got)
	}
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpNe, L: lit(int64(1)), R: lit(int64(2))}, Env{}); got != true {
		t.Fatalf("1!=2 = %v want true", got)
	}
}

func TestEval_DivModByZero(t *testing.T) {
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpDiv, L: lit(int64(1)), R: lit(int64(0))}, Env{}); got != nil {
		t.Fatalf("1/0 = %v want nil", got)
	}
	if got := evalOK(t, ast.BinaryExpr{Op: ast.OpMod, L: lit(int64(1)), R: lit(int64(0))}, Env{}); got != nil {
		t.Fatalf("1%%0 = %v want nil", got)
	}
}

func TestEval_Add_PropagatesError(t *testing.T) {
	_, errL := Eval(ast.BinaryExpr{Op: ast.OpAdd, L: ast.FuncCall{Name: "nope"}, R: lit(int64(1))}, Env{})
	if !errors.Is(errL, ErrUnknownFunc) {
		t.Fatalf("left: got %v want ErrUnknownFunc", errL)
	}
	_, errR := Eval(ast.BinaryExpr{Op: ast.OpAdd, L: lit(int64(1)), R: ast.FuncCall{Name: "nope"}}, Env{})
	if !errors.Is(errR, ErrUnknownFunc) {
		t.Fatalf("right: got %v want ErrUnknownFunc", errR)
	}
}

func TestEval_UnknownFunc(t *testing.T) {
	if _, err := Eval(ast.FuncCall{Name: "nope"}, Env{}); !errors.Is(err, ErrUnknownFunc) {
		t.Fatalf("got %v want ErrUnknownFunc", err)
	}
}

func TestEval_Call_ArgError(t *testing.T) {
	_, err := Eval(ast.FuncCall{Name: "length", Args: []ast.Expr{ast.FuncCall{Name: "nope"}}}, Env{})
	if !errors.Is(err, ErrUnknownFunc) {
		t.Fatalf("got %v want ErrUnknownFunc", err)
	}
}

// A nil expression — constructible only programmatically, never by the parser —
// reports ErrNilExpr instead of panicking.
func TestEvalNilExprErrors(t *testing.T) {
	_, err := Eval(nil, Env{})
	if !errors.Is(err, ErrNilExpr) {
		t.Errorf("got %v, want ErrNilExpr", err)
	}
}
