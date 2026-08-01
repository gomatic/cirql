package eval

import (
	"math"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

// The binary operators: equality, ordering and arithmetic. Split from the
// expression walker because the two change for different reasons — the walker
// grows when the AST grows a node shape, this file when the language grows an
// operator — and because each operator family is now a named subset of
// ast.BinOp whose switch must stay total over its own domain.

// apply evaluates a non-logical binary operator over two values.
func apply(op ast.BinOp, l, r value.Value) (value.Value, error) {
	if op == ast.OpEq {
		return value.Equal(l, r), nil
	}
	if op == ast.OpNe {
		return !value.Equal(l, r), nil
	}
	if op == ast.OpAdd {
		return value.Add(l, r)
	}
	if isCompare(op) {
		return compare(op, l, r)
	}
	return arith(op, l, r)
}

// isCompare reports whether op is an ordering comparison.
func isCompare(op ast.BinOp) bool {
	return op == ast.OpGt || op == ast.OpLt || op == ast.OpGe || op == ast.OpLe
}

// compare evaluates an ordering comparison.
func compare(op ast.BinOp, l, r value.Value) (value.Value, error) {
	c, err := value.Compare(l, r)
	if err != nil {
		return nil, err
	}
	return cmpResult(comparisonOp(op), ordering(c)), nil
}

// ordering is the -1/0/1 result of value.Compare.
type ordering int

// comparisonOp is the subset of ast.BinOp that orders two values. The narrow
// type is what lets cmpResult be total over its own domain: apply routes
// equality, logic and arithmetic elsewhere, so a switch over the whole of
// ast.BinOp here would carry six arms nothing can reach — and the default that
// stood in for them answered EVERY unrecognised operator as "<=", which is a
// silently wrong boolean rather than a refusal.
type comparisonOp ast.BinOp

const (
	cmpGreater      comparisonOp = comparisonOp(ast.OpGt)
	cmpLess         comparisonOp = comparisonOp(ast.OpLt)
	cmpGreaterEqual comparisonOp = comparisonOp(ast.OpGe)
	cmpLessEqual    comparisonOp = comparisonOp(ast.OpLe)
)

// cmpResult turns a -1/0/1 comparison into the boolean the operator wants.
// An operator outside the declared four is false: no ordering claim holds for
// an operator that makes no ordering claim.
func cmpResult(op comparisonOp, c ordering) bool {
	switch op {
	case cmpGreater:
		return int(c) > 0
	case cmpLess:
		return int(c) < 0
	case cmpGreaterEqual:
		return int(c) >= 0
	case cmpLessEqual:
		return int(c) <= 0
	default:
		return false
	}
}

// arith evaluates -, *, /, %. Subtraction and multiplication preserve int64
// when both operands are integers — matching + (value.Add) — so the numeric
// model is consistent across the additive and multiplicative operators;
// division and modulo use the double model, yielding null on a zero divisor.
func arith(op ast.BinOp, l, r value.Value) (value.Value, error) {
	if op == ast.OpSub || op == ast.OpMul {
		if v, ok := intArith(op, l, r); ok {
			return v, nil
		}
	}
	x, err := value.AsFloat(l)
	if err != nil {
		return nil, ErrType
	}
	y, err := value.AsFloat(r)
	if err != nil {
		return nil, ErrType
	}
	return arithFloat(arithmeticOp(op), operand(x), operand(y)), nil
}

// intArith computes an integer subtraction or multiplication when both operands
// are int64, reporting ok=false to fall through to the float model otherwise.
func intArith(op ast.BinOp, l, r value.Value) (value.Value, bool) {
	li, lok := l.(int64)
	ri, rok := r.(int64)
	if !lok || !rok {
		return nil, false
	}
	if op == ast.OpSub {
		return li - ri, true
	}
	return li * ri, true
}

// operand is one numeric operand of an arithmetic operation in the double model.
type operand float64

// arithmeticOp is the subset of ast.BinOp that combines two numbers. Addition
// is absent deliberately: apply answers it with value.Add, which preserves
// int64, so it never reaches the double model here.
type arithmeticOp ast.BinOp

const (
	arithSub arithmeticOp = arithmeticOp(ast.OpSub)
	arithMul arithmeticOp = arithmeticOp(ast.OpMul)
	arithDiv arithmeticOp = arithmeticOp(ast.OpDiv)
	arithMod arithmeticOp = arithmeticOp(ast.OpMod)
)

// arithFloat applies the numeric operator to two operands. An operator outside
// the declared four yields null — the same answer a zero divisor gives —
// rather than the modulus the default arm used to return for anything it did
// not recognise.
func arithFloat(op arithmeticOp, x, y operand) value.Value {
	switch op {
	case arithSub:
		return float64(x - y)
	case arithMul:
		return float64(x * y)
	case arithDiv:
		return zeroGuard(x/y, y)
	case arithMod:
		return zeroGuard(operand(math.Mod(float64(x), float64(y))), y)
	default:
		return nil
	}
}

// zeroGuard returns null when the divisor is zero, else the computed result.
func zeroGuard(result, divisor operand) value.Value {
	if divisor == 0 {
		return nil
	}
	return float64(result)
}

// evalCall dispatches a builtin function call.
func evalCall(n ast.FuncCall, env Env) (value.Value, error) {
	fn, ok := builtins[n.Name]
	if !ok {
		return nil, ErrUnknownFunc
	}
	args, err := evalArgs(n.Args, env)
	if err != nil {
		return nil, err
	}
	return fn(args, env)
}

// evalArgs evaluates every argument expression in order.
func evalArgs(exprs []ast.Expr, env Env) ([]value.Value, error) {
	args := make([]value.Value, 0, len(exprs))
	for _, e := range exprs {
		v, err := Eval(e, env)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	return args, nil
}
