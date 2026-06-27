package eval

import (
	"sort"
	"strings"

	"github.com/gomatic/cirql/value"
)

// builtin is a cirql builtin function.
type builtin func(args []value.Value, env Env) (value.Value, error)

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
	"now":        biNow,
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

// kindName is the spec type name for a kind.
func kindName(k value.Kind) string {
	switch k {
	case value.KindBool:
		return "bool"
	case value.KindInt, value.KindFloat:
		return "number"
	case value.KindString:
		return "string"
	case value.KindList:
		return "list"
	case value.KindObject:
		return "object"
	default:
		return "null"
	}
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

func biUpper(args []value.Value, _ Env) (value.Value, error) {
	return stringOp(args, strings.ToUpper)
}

func biLower(args []value.Value, _ Env) (value.Value, error) {
	return stringOp(args, strings.ToLower)
}

func biTrim(args []value.Value, _ Env) (value.Value, error) {
	return stringOp(args, strings.TrimSpace)
}

// stringOp applies a string transform to a single string argument.
func stringOp(args []value.Value, fn func(string) string) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	s, serr := value.AsString(x)
	if serr != nil {
		return nil, ErrType
	}
	return fn(s), nil
}

func biSplit(args []value.Value, _ Env) (value.Value, error) {
	s, sep, err := twoStrings(args)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(s, sep)
	out := make([]value.Value, len(parts))
	for i, p := range parts {
		out[i] = p
	}
	return out, nil
}

func biJoin(args []value.Value, _ Env) (value.Value, error) {
	list, sep, err := listAndSep(args)
	if err != nil {
		return nil, err
	}
	parts := make([]string, 0, len(list))
	for _, item := range list {
		s, serr := value.AsString(item)
		if serr != nil {
			return nil, ErrType
		}
		parts = append(parts, s)
	}
	return strings.Join(parts, sep), nil
}

// listAndSep extracts a (list, separator-string) argument pair.
func listAndSep(args []value.Value) ([]value.Value, string, error) {
	a, b, err := arg2(args)
	if err != nil {
		return nil, "", err
	}
	list, lerr := value.AsList(a)
	sep, serr := value.AsString(b)
	if lerr != nil || serr != nil {
		return nil, "", ErrType
	}
	return list, sep, nil
}

// twoStrings extracts two string arguments.
func twoStrings(args []value.Value) (string, string, error) {
	a, b, err := arg2(args)
	if err != nil {
		return "", "", err
	}
	as, aerr := value.AsString(a)
	bs, berr := value.AsString(b)
	if aerr != nil || berr != nil {
		return "", "", ErrType
	}
	return as, bs, nil
}

func biContains(args []value.Value, _ Env) (value.Value, error) {
	a, b, err := arg2(args)
	if err != nil {
		return nil, err
	}
	if s, ok := a.(string); ok {
		sub, serr := value.AsString(b)
		if serr != nil {
			return nil, ErrType
		}
		return strings.Contains(s, sub), nil
	}
	list, lerr := value.AsList(a)
	if lerr != nil {
		return nil, ErrType
	}
	return listContains(list, b), nil
}

// listContains reports whether list has an element equal to x.
func listContains(list []value.Value, x value.Value) bool {
	for _, item := range list {
		if value.Equal(item, x) {
			return true
		}
	}
	return false
}

func biStartsWith(args []value.Value, _ Env) (value.Value, error) {
	s, prefix, err := twoStrings(args)
	if err != nil {
		return nil, err
	}
	return strings.HasPrefix(s, prefix), nil
}

func biNow(args []value.Value, env Env) (value.Value, error) {
	if len(args) != 0 {
		return nil, ErrArity
	}
	if env.Now == nil {
		return int64(0), nil
	}
	return env.Now(), nil
}

func biFlatten(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	list, lerr := value.AsList(x)
	if lerr != nil {
		return nil, ErrType
	}
	return flattenOne(list), nil
}

// flattenOne flattens one level of nested lists.
func flattenOne(list []value.Value) []value.Value {
	out := make([]value.Value, 0, len(list))
	for _, item := range list {
		if inner, ok := item.([]value.Value); ok {
			out = append(out, inner...)
			continue
		}
		out = append(out, item)
	}
	return out
}

func biDistinct(args []value.Value, _ Env) (value.Value, error) {
	x, err := arg1(args)
	if err != nil {
		return nil, err
	}
	list, lerr := value.AsList(x)
	if lerr != nil {
		return nil, ErrType
	}
	return distinctList(list), nil
}

// distinctList returns the list with duplicate (value-equal) elements removed.
func distinctList(list []value.Value) []value.Value {
	out := make([]value.Value, 0, len(list))
	for _, item := range list {
		if !listContains(out, item) {
			out = append(out, item)
		}
	}
	return out
}

func biCoalesce(args []value.Value, _ Env) (value.Value, error) {
	if len(args) == 0 {
		return nil, ErrArity
	}
	for _, a := range args {
		if a != nil {
			return a, nil
		}
	}
	return nil, nil
}
