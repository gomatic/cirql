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

// TestArg2RequiresExactlyTwoArguments names arg2's claim. Every two-argument
// builtin destructures through it, so an off-by-one here is an index panic
// inside a builtin on user-supplied query text. "Exactly" means both
// directions: one argument too few and one too many are equally rejected, and
// the rejection is ErrArity so a caller can tell a misused function from a type
// error.
func TestArg2RequiresExactlyTwoArguments(t *testing.T) {
	for _, args := range [][]value.Value{
		nil,
		{int64(1)},
		{int64(1), int64(2), int64(3)},
	} {
		if _, _, err := arg2(args); !errors.Is(err, ErrArity) {
			t.Fatalf("arg2(%d args) = %v, want ErrArity", len(args), err)
		}
	}

	a, b, err := arg2([]value.Value{int64(1), "two"})
	if err != nil {
		t.Fatalf("arg2 of exactly two arguments: %v", err)
	}
	if a != value.Value(int64(1)) || b != value.Value("two") {
		t.Fatalf("arg2 returned (%v, %v), want the arguments in order", a, b)
	}
}

// TestKindNameCoversEveryDeclaredKind pins the table's totality. The `type`
// builtin surfaces these names to the user, so a kind falling through to the
// fallback would report a list or an object as "null" — a lie the query author
// would then branch on.
func TestKindNameCoversEveryDeclaredKind(t *testing.T) {
	for kind, want := range map[value.Kind]string{
		value.KindNull:   typeNull,
		value.KindBool:   typeBool,
		value.KindInt:    typeNumber,
		value.KindFloat:  typeNumber,
		value.KindString: typeString,
		value.KindList:   typeList,
		value.KindObject: typeObject,
	} {
		if got := kindName(kind); got != want {
			t.Fatalf("kindName(%v) = %q, want %q", kind, got, want)
		}
	}
}

// TestKindNameFallsBackToNullForAnUndeclaredKind covers the fallback, which is
// reachable only by constructing a Kind the package does not declare — the same
// answer value.KindOf gives for an unrecognized concrete type.
func TestKindNameFallsBackToNullForAnUndeclaredKind(t *testing.T) {
	if got := kindName(value.Kind(200)); got != typeNull {
		t.Fatalf("kindName(undeclared) = %q, want %q", got, typeNull)
	}
}
