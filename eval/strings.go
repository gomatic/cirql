package eval

import (
	"strings"

	value "github.com/gomatic/go-json"
)

// The string builtins: upper, lower, trim, split, join, contains, startsWith,
// and now. Split out from the registry and the collection builtins so each file
// holds one family — builtins.go stays the single place a name is bound, and
// adding a string function does not grow the file that owns the dispatch.

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
