// Package eval evaluates cirql expressions against an environment: the current
// object (for field access), the variable bindings (for $var), and an injected
// clock (for now()). Field access propagates null and arithmetic follows jq's
// double model; division or modulo by zero yields null.
package eval

import (
	"math"

	"github.com/gomatic/cirql/ast"
	value "github.com/gomatic/go-json"
)

// Env is the evaluation context for one expression.
type Env struct {
	Obj  value.Value
	Vars map[string]value.Value
	Now  func() int64 // epoch seconds; nil falls back to a zero clock
}

// Eval evaluates e against env.
func Eval(e ast.Expr, env Env) (value.Value, error) {
	switch n := e.(type) {
	case ast.Literal:
		return n.V, nil
	case ast.FieldAccess:
		return evalPath(env.Obj, n.Path), nil
	case ast.VarRef:
		return env.Vars[n.Name], nil
	case ast.UnaryExpr:
		return evalUnary(n, env)
	case ast.FuncCall:
		return evalCall(n, env)
	default:
		return evalBinary(e.(ast.BinaryExpr), env)
	}
}

// evalPath walks a field-access path, propagating null and mapping [] over lists.
func evalPath(cur value.Value, segs []ast.PathSegment) value.Value {
	if len(segs) == 0 {
		return cur
	}
	if segs[0].Iter {
		return evalIter(cur, segs[1:])
	}
	return evalField(cur, segs[0].Name, segs[1:])
}

// evalIter applies the remaining path to each element of a list; a non-list is
// null.
func evalIter(cur value.Value, rest []ast.PathSegment) value.Value {
	list, err := value.AsList(cur)
	if err != nil {
		return nil
	}
	out := make([]value.Value, 0, len(list))
	for _, item := range list {
		out = append(out, evalPath(item, rest))
	}
	return out
}

// evalField indexes one object field (missing or non-object yields null) and
// continues down the path.
func evalField(cur value.Value, name string, rest []ast.PathSegment) value.Value {
	obj, err := value.AsObject(cur)
	if err != nil {
		return nil
	}
	return evalPath(obj[name], rest)
}

// evalUnary evaluates a logical-not or arithmetic-negation.
func evalUnary(n ast.UnaryExpr, env Env) (value.Value, error) {
	x, err := Eval(n.X, env)
	if err != nil {
		return nil, err
	}
	if n.Op == ast.OpNot {
		return !value.Truthy(x), nil
	}
	return negate(x)
}

// negate arithmetically negates a number.
func negate(x value.Value) (value.Value, error) {
	if i, ok := x.(int64); ok {
		return -i, nil
	}
	f, err := value.AsFloat(x)
	if err != nil {
		return nil, ErrType
	}
	return -f, nil
}

// evalBinary evaluates a binary expression, short-circuiting logical operators.
func evalBinary(n ast.BinaryExpr, env Env) (value.Value, error) {
	if n.Op == ast.OpAnd || n.Op == ast.OpOr {
		return evalLogical(n, env)
	}
	l, err := Eval(n.L, env)
	if err != nil {
		return nil, err
	}
	r, err := Eval(n.R, env)
	if err != nil {
		return nil, err
	}
	return apply(n.Op, l, r)
}

// evalLogical evaluates && / || with short-circuit semantics.
func evalLogical(n ast.BinaryExpr, env Env) (value.Value, error) {
	l, err := Eval(n.L, env)
	if err != nil {
		return nil, err
	}
	if short, ok := shortCircuit(n.Op, value.Truthy(l)); ok {
		return short, nil
	}
	r, err := Eval(n.R, env)
	if err != nil {
		return nil, err
	}
	return value.Truthy(r), nil
}

// shortCircuit reports the determined result of a logical op from its left
// operand, or ok=false when the right operand must be evaluated.
func shortCircuit(op ast.BinOp, left bool) (result, ok bool) {
	if op == ast.OpAnd && !left {
		return false, true
	}
	if op == ast.OpOr && left {
		return true, true
	}
	return false, false
}

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
	return cmpResult(op, c), nil
}

// cmpResult turns a -1/0/1 comparison into the boolean the operator wants.
func cmpResult(op ast.BinOp, c int) bool {
	switch op {
	case ast.OpGt:
		return c > 0
	case ast.OpLt:
		return c < 0
	case ast.OpGe:
		return c >= 0
	default:
		return c <= 0
	}
}

// arith evaluates -, *, /, % in the double model; / and % by zero yield null.
func arith(op ast.BinOp, l, r value.Value) (value.Value, error) {
	x, err := value.AsFloat(l)
	if err != nil {
		return nil, ErrType
	}
	y, err := value.AsFloat(r)
	if err != nil {
		return nil, ErrType
	}
	return arithFloat(op, x, y), nil
}

// arithFloat applies the numeric operator to two floats.
func arithFloat(op ast.BinOp, x, y float64) value.Value {
	switch op {
	case ast.OpSub:
		return x - y
	case ast.OpMul:
		return x * y
	case ast.OpDiv:
		return zeroGuard(x/y, y)
	default:
		return zeroGuard(math.Mod(x, y), y)
	}
}

// zeroGuard returns null when the divisor is zero, else the computed result.
func zeroGuard(result, divisor float64) value.Value {
	if divisor == 0 {
		return nil
	}
	return result
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
