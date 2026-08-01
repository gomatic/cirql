// Package eval evaluates cirql expressions against an environment: the current
// object (for field access), the variable bindings (for $var), and an injected
// clock (for now()). Field access propagates null and arithmetic follows jq's
// double model; division or modulo by zero yields null.
package eval

import (
	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

// Env is the evaluation context for one expression.
type Env struct {
	Obj  value.Value
	Vars map[string]value.Value
	Now  func() int64 // epoch seconds; nil falls back to a zero clock
}

// Eval evaluates e against env. A nil expression — possible only in a
// programmatically assembled pipeline, never from the parser — reports
// ErrNilExpr rather than panicking.
func Eval(e ast.Expr, env Env) (value.Value, error) {
	if e == nil {
		return nil, ErrNilExpr
	}
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
	if segs[0].IsIter {
		return evalIter(cur, segs[1:])
	}
	return evalField(cur, fieldName(segs[0].Name), segs[1:])
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

// fieldName is the name of an object field addressed by one path segment.
type fieldName string

// evalField indexes one object field (missing or non-object yields null) and
// continues down the path.
func evalField(cur value.Value, name fieldName, rest []ast.PathSegment) value.Value {
	obj, err := value.AsObject(cur)
	if err != nil {
		return nil
	}
	return evalPath(obj[string(name)], rest)
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
	if short, ok := shortCircuit(n.Op, leftTruthy(value.Truthy(l))); ok {
		return short, nil
	}
	r, err := Eval(n.R, env)
	if err != nil {
		return nil, err
	}
	return value.Truthy(r), nil
}

// leftTruthy is the truthiness of the already-evaluated left operand of a logical operator.
type leftTruthy bool

// shortCircuit reports the determined isResult of a logical op from its isLeft
// operand, or isOk=false when the right operand must be evaluated.
func shortCircuit(op ast.BinOp, isLeft leftTruthy) (isResult, isOk bool) {
	if op == ast.OpAnd && !bool(isLeft) {
		return false, true
	}
	if op == ast.OpOr && bool(isLeft) {
		return true, true
	}
	return false, false
}
