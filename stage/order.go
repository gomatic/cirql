package stage

import (
	"slices"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
)

// The stages that reorder or shrink a result set without changing any element:
// sort, limit and uniq. Their contracts are about WHICH objects survive and in
// what order, never about what an object contains.

// sortExec reorders the result set by a key expression.
type sortExec struct {
	key    ast.Expr
	now    Clock
	isDesc bool
}

// keyedItem pairs an item with its precomputed sort key.
type keyedItem struct {
	key  value.Value
	item value.Value
}

func (s sortExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	keyed, err := s.keyed(in)
	if err != nil {
		return nil, err
	}
	var cmpErr error
	slices.SortStableFunc(keyed, func(a, b keyedItem) int {
		c, err := s.compare(a, b)
		if err != nil {
			cmpErr = err
		}
		return c
	})
	if cmpErr != nil {
		return nil, cmpErr
	}
	return unkey(keyed), nil
}

func (s sortExec) keyed(in pipeline.ResultSet) ([]keyedItem, error) {
	out := make([]keyedItem, 0, len(in))
	for _, item := range in {
		k, err := eval.Eval(s.key, envFor(item, s.now))
		if err != nil {
			return nil, err
		}
		out = append(out, keyedItem{key: k, item: item})
	}
	return out, nil
}

// compare orders two keyed items, or reports that the keys are incomparable.
func (s sortExec) compare(a, b keyedItem) (int, error) {
	c, err := value.Compare(a.key, b.key)
	if err != nil {
		return 0, err
	}
	if s.isDesc {
		return -c, nil
	}
	return c, nil
}

// unkey drops the sort keys, returning the ordered items.
func unkey(keyed []keyedItem) pipeline.ResultSet {
	out := make(pipeline.ResultSet, len(keyed))
	for i, k := range keyed {
		out[i] = k.item
	}
	return out
}

// limitExec truncates the result set to at most n elements. A negative n —
// possible only in a programmatically assembled pipeline, never from the
// parser — behaves as zero.
type limitExec struct{ n int }

func (s limitExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	n := max(s.n, 0)
	if n < len(in) {
		return in[:n], nil
	}
	return in, nil
}

// uniqExec deduplicates the result set, by key when set else whole-value.
type uniqExec struct {
	key ast.Expr
	now Clock
}

func (s uniqExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	out := pipeline.ResultSet{}
	seen := []value.Value{}
	for _, item := range in {
		k, err := s.keyOf(item)
		if err != nil {
			return nil, err
		}
		if !containsValue(seen, k) {
			seen = append(seen, k)
			out = append(out, item)
		}
	}
	return out, nil
}

func (s uniqExec) keyOf(item value.Value) (value.Value, error) {
	if s.key == nil {
		return item, nil
	}
	return eval.Eval(s.key, envFor(item, s.now))
}

// containsValue reports whether seen holds a value equal to k.
func containsValue(seen []value.Value, k value.Value) bool {
	for _, v := range seen {
		if value.Equal(v, k) {
			return true
		}
	}
	return false
}
