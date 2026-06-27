// Package pipeline holds the cirql result-set model and the sequential stage
// runner. A ResultSet is the list of values flowing between stages; Normalize
// turns an arbitrary decoded JSON value into the initial result set per the
// source-normalization rules (spec §5.3).
package pipeline

import "github.com/gomatic/cirql/value"

// ResultSet is the data flowing between pipeline stages.
type ResultSet = []value.Value

// Stage is one executable step of a pipeline.
type Stage interface {
	Execute(in ResultSet) (ResultSet, error)
}

// Normalize turns a decoded JSON value into the initial result set: a list maps
// element-by-element (non-object elements wrapped as {"value": elem}); any other
// value becomes a single-element set (also wrapped when not an object).
func Normalize(v value.Value) ResultSet {
	if list, ok := v.([]value.Value); ok {
		return normalizeList(list)
	}
	return ResultSet{wrapNonObject(v)}
}

// normalizeList normalizes each element of a list.
func normalizeList(list []value.Value) ResultSet {
	out := make(ResultSet, len(list))
	for i, item := range list {
		out[i] = wrapNonObject(item)
	}
	return out
}

// wrapNonObject returns objects unchanged and wraps any other value as
// {"value": v} so downstream stages always see objects.
func wrapNonObject(v value.Value) value.Value {
	if _, ok := v.(map[string]value.Value); ok {
		return v
	}
	return map[string]value.Value{"value": v}
}

// RunStages threads in through every stage left to right, stopping at the first
// stage error.
func RunStages(stages []Stage, in ResultSet) (ResultSet, error) {
	cur := in
	for _, s := range stages {
		next, err := s.Execute(cur)
		if err != nil {
			return nil, err
		}
		cur = next
	}
	return cur, nil
}
