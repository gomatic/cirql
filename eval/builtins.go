package eval

import (
	"sort"

	value "github.com/gomatic/go-json"
)

// builtin is a cirql builtin function.
type builtin func(args []value.Value, env Env) (value.Value, error)

// Spec type names returned by the type() builtin (spec §5.5).
const (
	typeNull   = "null"
	typeBool   = "bool"
	typeNumber = "number"
	typeString = "string"
	typeList   = "list"
	typeObject = "object"
)

// nameNow is the registry key of the now() builtin.
const nameNow = "now"

// builtins is the registry of callable builtins (spec §5.5).
var builtins = map[string]builtin{
	"length":     biLength,
	"keys":       biKeys,
	"values":     biValues,
	"type":       biType,
	"toInt":      biToInt,
	"toFloat":    biToFloat,
	"toString":   biToString,
	"upper":      biUpper,
	"lower":      biLower,
	"trim":       biTrim,
	"split":      biSplit,
	"join":       biJoin,
	"contains":   biContains,
	"startsWith": biStartsWith,
	nameNow:      biNow,
	"flatten":    biFlatten,
	"distinct":   biDistinct,
	"coalesce":   biCoalesce,
}

// arg1 extracts the single argument or reports an arity error.
func arg1(args []value.Value) (value.Value, error) {
	if len(args) != 1 {
		return nil, ErrArity
	}
	return args[0], nil
}

// arg2 extracts exactly two arguments or reports an arity error.
func arg2(args []value.Value) (a, b value.Value, err error) {
	if len(args) != 2 {
		return nil, nil, ErrArity
	}
	return args[0], args[1], nil
}

func biLength(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	switch v := x.(type) {
	case string:
		return int64(len([]rune(v))), nil
	case []value.Value:
		return int64(len(v)), nil
	case map[string]value.Value:
		return int64(len(v)), nil
	default:
		return nil, ErrType
	}
}

func biKeys(args []value.Value, _ Env) (value.Value, error) {
	obj, err := unaryObject(args)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]value.Value, len(keys))
	for i, k := range keys {
		out[i] = k
	}
	return out, nil
}

func biValues(args []value.Value, _ Env) (value.Value, error) {
	obj, err := unaryObject(args)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]value.Value, len(keys))
	for i, k := range keys {
		out[i] = obj[k]
	}
	return out, nil
}

// unaryObject extracts a single object argument.
func unaryObject(args []value.Value) (map[string]value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	obj, oerr := value.AsObject(x)
	if oerr != nil {
		return nil, ErrType
	}
	return obj, nil
}

func biType(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	return kindName(value.KindOf(x)), nil
}

// kindNames is the spec type name for every kind. A table rather than a switch
// so that the mapping is total by construction: a Kind added to go-json without
// a name here is a build-time omission an exhaustiveness check can see, not a
// value that silently reports itself as null.
var kindNames = map[value.Kind]string{
	value.KindNull:   typeNull,
	value.KindBool:   typeBool,
	value.KindInt:    typeNumber,
	value.KindFloat:  typeNumber,
	value.KindString: typeString,
	value.KindList:   typeList,
	value.KindObject: typeObject,
}

// kindName is the spec type name for a kind. A Kind outside the declared set
// cannot arrive from value.KindOf, but the language permits one to be
// constructed, and the spec has no name for it — null is the same answer
// KindOf gives an unrecognized concrete type.
func kindName(k value.Kind) string {
	if name, ok := kindNames[k]; ok {
		return name
	}
	return typeNull
}

func biToInt(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	n, cerr := value.AsInt(x)
	if cerr != nil {
		return nil, ErrType
	}
	return n, nil
}

func biToFloat(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	f, cerr := value.AsFloat(x)
	if cerr != nil {
		return nil, ErrType
	}
	return f, nil
}

func biToString(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	s, _ := value.Add("", x) // concat with "" string-coerces a scalar
	if str, ok := s.(string); ok {
		return str, nil
	}
	return nil, ErrType
}
