package eval

import value "github.com/gomatic/go-json"

// The collection builtins: flatten, distinct, coalesce — the ones whose
// argument is a list and whose answer depends on comparing elements rather than
// on formatting them.

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
