package eval

import (
	"errors"
	"reflect"
	"testing"

	value "github.com/gomatic/go-json"
)

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
