package eval

import (
	"errors"
	"reflect"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

func TestBuiltin_StringOps(t *testing.T) {
	if v := callOK(t, "upper", "ab"); v != "AB" {
		t.Fatalf("upper=%v", v)
	}
	if v := callOK(t, "lower", "AB"); v != "ab" {
		t.Fatalf("lower=%v", v)
	}
	if v := callOK(t, "trim", "  x  "); v != "x" {
		t.Fatalf("trim=%v", v)
	}
	if _, err := call(t, "upper", int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("upper(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "lower"); !errors.Is(err, ErrArity) {
		t.Fatalf("lower() err=%v want ErrArity", err)
	}
}

func TestBuiltin_SplitJoin(t *testing.T) {
	got := callOK(t, "split", "a,b,c", ",")
	if !reflect.DeepEqual(got, []value.Value{"a", "b", "c"}) {
		t.Fatalf("split=%#v", got)
	}
	if v := callOK(t, "join", []value.Value{"a", "b"}, "-"); v != "a-b" {
		t.Fatalf("join=%v want a-b", v)
	}
	if _, err := call(t, "split", int64(1), ","); !errors.Is(err, ErrType) {
		t.Fatalf("split(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "split", "a"); !errors.Is(err, ErrArity) {
		t.Fatalf("split arity err=%v want ErrArity", err)
	}
	if _, err := call(t, "join", int64(1), "-"); !errors.Is(err, ErrType) {
		t.Fatalf("join(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "join", []value.Value{int64(1)}, "-"); !errors.Is(err, ErrType) {
		t.Fatalf("join(non-string elems) err=%v want ErrType", err)
	}
	if _, err := call(t, "join", []value.Value{}); !errors.Is(err, ErrArity) {
		t.Fatalf("join arity err=%v want ErrArity", err)
	}
}

func TestBuiltin_Contains(t *testing.T) {
	if v := callOK(t, "contains", "hello", "ell"); v != true {
		t.Fatalf("contains(str)=%v want true", v)
	}
	if v := callOK(t, "contains", []value.Value{"a", "b"}, "b"); v != true {
		t.Fatalf("contains(list)=%v want true", v)
	}
	if v := callOK(t, "contains", []value.Value{"a"}, "z"); v != false {
		t.Fatalf("contains(list miss)=%v want false", v)
	}
	if _, err := call(t, "contains", "x", int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("contains(str,int) err=%v want ErrType", err)
	}
	if _, err := call(t, "contains", int64(1), int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("contains(int,int) err=%v want ErrType", err)
	}
	if _, err := call(t, "contains", "x"); !errors.Is(err, ErrArity) {
		t.Fatalf("contains arity err=%v want ErrArity", err)
	}
}

func TestBuiltin_StartsWith(t *testing.T) {
	if v := callOK(t, "startsWith", "abc", "ab"); v != true {
		t.Fatalf("startsWith=%v want true", v)
	}
	if _, err := call(t, "startsWith", "a"); !errors.Is(err, ErrArity) {
		t.Fatalf("startsWith arity err=%v want ErrArity", err)
	}
}

func TestBuiltin_Now(t *testing.T) {
	v, err := Eval(ast.FuncCall{Name: "now"}, Env{Now: func() int64 { return 1234 }})
	if err != nil || v != int64(1234) {
		t.Fatalf("now()=%v,%v want 1234", v, err)
	}
	zero, err := Eval(ast.FuncCall{Name: "now"}, Env{})
	if err != nil || zero != int64(0) {
		t.Fatalf("now() zero clock=%v,%v want 0", zero, err)
	}
	if _, err := Eval(ast.FuncCall{Name: "now", Args: []ast.Expr{lit(int64(1))}}, Env{}); !errors.Is(err, ErrArity) {
		t.Fatalf("now(x) err=%v want ErrArity", err)
	}
}
