package eval

import (
	"errors"
	"reflect"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
)

func call(t *testing.T, name string, args ...value.Value) (value.Value, error) {
	t.Helper()
	exprs := make([]ast.Expr, len(args))
	for i, a := range args {
		exprs[i] = ast.Literal{V: a}
	}
	return Eval(ast.FuncCall{Name: name, Args: exprs}, Env{Vars: map[string]value.Value{}})
}

func callOK(t *testing.T, name string, args ...value.Value) value.Value {
	t.Helper()
	v, err := call(t, name, args...)
	if err != nil {
		t.Fatalf("%s error: %v", name, err)
	}
	return v
}

func TestBuiltin_Length(t *testing.T) {
	if v := callOK(t, "length", "abc"); v != int64(3) {
		t.Fatalf("length(abc)=%v want 3", v)
	}
	if v := callOK(t, "length", []value.Value{int64(1), int64(2)}); v != int64(2) {
		t.Fatalf("length(list)=%v want 2", v)
	}
	if v := callOK(t, "length", map[string]value.Value{"a": int64(1)}); v != int64(1) {
		t.Fatalf("length(obj)=%v want 1", v)
	}
	if _, err := call(t, "length", int64(5)); !errors.Is(err, ErrType) {
		t.Fatalf("length(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "length"); !errors.Is(err, ErrArity) {
		t.Fatalf("length() err=%v want ErrArity", err)
	}
}

func TestBuiltin_KeysValues(t *testing.T) {
	obj := map[string]value.Value{"b": int64(2), "a": int64(1)}
	keys := callOK(t, "keys", obj)
	if !reflect.DeepEqual(keys, []value.Value{"a", "b"}) {
		t.Fatalf("keys=%#v want [a b]", keys)
	}
	vals := callOK(t, "values", obj)
	if !reflect.DeepEqual(vals, []value.Value{int64(1), int64(2)}) {
		t.Fatalf("values=%#v want [1 2]", vals)
	}
	if _, err := call(t, "keys", int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("keys(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "values"); !errors.Is(err, ErrArity) {
		t.Fatalf("values() err=%v want ErrArity", err)
	}
	if _, err := call(t, "keys"); !errors.Is(err, ErrArity) {
		t.Fatalf("keys() err=%v want ErrArity", err)
	}
}

func TestBuiltin_Type(t *testing.T) {
	cases := map[string]value.Value{
		"null": nil, "bool": true, "number": int64(1),
		"string": "s", "list": []value.Value{}, "object": map[string]value.Value{},
	}
	for want, v := range cases {
		if got := callOK(t, "type", v); got != want {
			t.Fatalf("type(%#v)=%v want %v", v, got, want)
		}
	}
	if got := callOK(t, "type", 1.5); got != "number" {
		t.Fatalf("type(float)=%v want number", got)
	}
	if _, err := call(t, "type"); !errors.Is(err, ErrArity) {
		t.Fatalf("type() err=%v want ErrArity", err)
	}
}

func TestBuiltin_Conversions(t *testing.T) {
	if v := callOK(t, "toInt", 3.9); v != int64(3) {
		t.Fatalf("toInt(3.9)=%v want 3", v)
	}
	if v := callOK(t, "toFloat", int64(4)); v != 4.0 {
		t.Fatalf("toFloat(4)=%v want 4.0", v)
	}
	if v := callOK(t, "toString", int64(7)); v != "7" {
		t.Fatalf("toString(7)=%v want 7", v)
	}
	if _, err := call(t, "toInt", "x"); !errors.Is(err, ErrType) {
		t.Fatalf("toInt(x) err=%v want ErrType", err)
	}
	if _, err := call(t, "toFloat", "x"); !errors.Is(err, ErrType) {
		t.Fatalf("toFloat(x) err=%v want ErrType", err)
	}
	if _, err := call(t, "toString", []value.Value{}); !errors.Is(err, ErrType) {
		t.Fatalf("toString(list) err=%v want ErrType", err)
	}
	for _, name := range []string{"toInt", "toFloat", "toString"} {
		if _, err := call(t, name); !errors.Is(err, ErrArity) {
			t.Fatalf("%s() err=%v want ErrArity", name, err)
		}
	}
}

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

func TestBuiltin_Flatten(t *testing.T) {
	in := []value.Value{[]value.Value{int64(1), int64(2)}, int64(3)}
	got := callOK(t, "flatten", in)
	if !reflect.DeepEqual(got, []value.Value{int64(1), int64(2), int64(3)}) {
		t.Fatalf("flatten=%#v", got)
	}
	if _, err := call(t, "flatten", int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("flatten(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "flatten"); !errors.Is(err, ErrArity) {
		t.Fatalf("flatten arity err=%v want ErrArity", err)
	}
}

func TestBuiltin_Distinct(t *testing.T) {
	in := []value.Value{int64(1), int64(1), int64(2)}
	got := callOK(t, "distinct", in)
	if !reflect.DeepEqual(got, []value.Value{int64(1), int64(2)}) {
		t.Fatalf("distinct=%#v", got)
	}
	if _, err := call(t, "distinct", int64(1)); !errors.Is(err, ErrType) {
		t.Fatalf("distinct(int) err=%v want ErrType", err)
	}
	if _, err := call(t, "distinct"); !errors.Is(err, ErrArity) {
		t.Fatalf("distinct arity err=%v want ErrArity", err)
	}
}

func TestBuiltin_Coalesce(t *testing.T) {
	if v := callOK(t, "coalesce", nil, "x"); v != "x" {
		t.Fatalf("coalesce=%v want x", v)
	}
	if v := callOK(t, "coalesce", nil, nil); v != nil {
		t.Fatalf("coalesce all nil=%v want nil", v)
	}
	if _, err := call(t, "coalesce"); !errors.Is(err, ErrArity) {
		t.Fatalf("coalesce arity err=%v want ErrArity", err)
	}
}
