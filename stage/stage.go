// Package stage implements the cirql transform stages and Build, which maps a
// parsed stage node to its executor. Source stages (file/http/query) report
// ErrStageUnsupported until the sources sub-project provides them; the stdin
// source is the identity stage (its input is the pipeline's input).
package stage

import (
	"fmt"
	"slices"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
	"github.com/gomatic/cirql/value"
)

// Clock supplies epoch-seconds for the now() builtin; nil yields a zero clock.
type Clock = func() int64

// Build maps a parsed stage to its executor. A source stage other than stdin
// reports ErrStageUnsupported.
func Build(s ast.Stage, now Clock) (pipeline.Stage, error) {
	switch n := s.(type) {
	case ast.MapStage:
		return mapExec{mappings: n.Mappings, now: now}, nil
	case ast.FlatMapStage:
		return flatMapExec{mappings: n.Mappings, now: now}, nil
	case ast.FilterStage:
		return filterExec{cond: n.Cond, now: now}, nil
	case ast.ReduceStage:
		return reduceExec{op: n.Op, arg: n.Arg, now: now}, nil
	case ast.SortStage:
		return sortExec{key: n.Key, desc: n.Desc, now: now}, nil
	case ast.LimitStage:
		return limitExec{n: n.N}, nil
	case ast.UniqStage:
		return uniqExec{key: n.Key, now: now}, nil
	case ast.StdinStage:
		return stdinExec{}, nil
	default:
		return nil, ErrStageUnsupported
	}
}

// envFor builds the evaluation environment for one object.
func envFor(item value.Value, now Clock) eval.Env {
	return eval.Env{Obj: item, Now: now}
}

// stdinExec yields its input unchanged: the pipeline input is the stdin source.
type stdinExec struct{}

func (stdinExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	return in, nil
}

// mapExec transforms each object 1:1.
type mapExec struct {
	mappings []ast.Mapping
	now      Clock
}

func (s mapExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	out := make(pipeline.ResultSet, 0, len(in))
	for _, item := range in {
		obj, err := s.one(item)
		if err != nil {
			return nil, err
		}
		out = append(out, obj)
	}
	return out, nil
}

func (s mapExec) one(item value.Value) (value.Value, error) {
	res := make(map[string]value.Value, len(s.mappings))
	for _, m := range s.mappings {
		v, err := eval.Eval(m.Expr, envFor(item, s.now))
		if err != nil {
			return nil, err
		}
		res[m.Key] = v
	}
	return res, nil
}

// flatMapExec transforms and flattens: list-valued mappings expand into one
// output object per element.
type flatMapExec struct {
	mappings []ast.Mapping
	now      Clock
}

// kv is one evaluated mapping.
type kv struct {
	key string
	val value.Value
}

func (s flatMapExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	out := pipeline.ResultSet{}
	for _, item := range in {
		rows, err := s.expand(item)
		if err != nil {
			return nil, err
		}
		out = append(out, rows...)
	}
	return out, nil
}

func (s flatMapExec) expand(item value.Value) ([]value.Value, error) {
	ev, err := s.evalAll(item)
	if err != nil {
		return nil, err
	}
	n := fanout(ev)
	rows := make([]value.Value, n)
	for i := range rows {
		rows[i] = rowAt(ev, i)
	}
	return rows, nil
}

func (s flatMapExec) evalAll(item value.Value) ([]kv, error) {
	ev := make([]kv, len(s.mappings))
	for i, m := range s.mappings {
		v, err := eval.Eval(m.Expr, envFor(item, s.now))
		if err != nil {
			return nil, err
		}
		ev[i] = kv{key: m.Key, val: v}
	}
	return ev, nil
}

// fanout is the number of output rows: the longest list-valued mapping, min 1.
func fanout(ev []kv) int {
	n := 1
	for _, e := range ev {
		if l, ok := e.val.([]value.Value); ok && len(l) > n {
			n = len(l)
		}
	}
	return n
}

// rowAt builds the i-th expanded row: list mappings take element i, scalars repeat.
func rowAt(ev []kv, i int) value.Value {
	row := make(map[string]value.Value, len(ev))
	for _, e := range ev {
		row[e.key] = elemAt(e.val, i)
	}
	return row
}

// elemAt returns the i-th element of a list (nil past the end) or a scalar as-is.
func elemAt(v value.Value, i int) value.Value {
	l, ok := v.([]value.Value)
	if !ok {
		return v
	}
	if i < len(l) {
		return l[i]
	}
	return nil
}

// filterExec keeps objects whose condition is truthy.
type filterExec struct {
	cond ast.Expr
	now  Clock
}

func (s filterExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	out := pipeline.ResultSet{}
	for _, item := range in {
		v, err := eval.Eval(s.cond, envFor(item, s.now))
		if err != nil {
			return nil, err
		}
		if value.Truthy(v) {
			out = append(out, item)
		}
	}
	return out, nil
}

// reduceExec aggregates the result set into a single-element set.
type reduceExec struct {
	op  ast.ReduceOp
	arg ast.Expr
	now Clock
}

func (s reduceExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	v, err := s.reduce(in)
	if err != nil {
		return nil, err
	}
	return pipeline.ResultSet{v}, nil
}

func (s reduceExec) reduce(in pipeline.ResultSet) (value.Value, error) {
	switch s.op {
	case ast.OpCount:
		return int64(len(in)), nil
	case ast.OpFirst:
		return firstOrNil(in), nil
	case ast.OpLast:
		return lastOrNil(in), nil
	case ast.OpCollect:
		return s.collect(in)
	case ast.OpGroupBy:
		return s.groupBy(in)
	default:
		return s.aggregate(in)
	}
}

// valueOf is the value a reduce works on for one object: the arg expression, or
// the object itself when no arg is given.
func (s reduceExec) valueOf(item value.Value) (value.Value, error) {
	if s.arg == nil {
		return item, nil
	}
	return eval.Eval(s.arg, envFor(item, s.now))
}

func (s reduceExec) collect(in pipeline.ResultSet) (value.Value, error) {
	out := make([]value.Value, 0, len(in))
	for _, item := range in {
		v, err := s.valueOf(item)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

func (s reduceExec) groupBy(in pipeline.ResultSet) (value.Value, error) {
	groups := map[string]value.Value{}
	for _, item := range in {
		v, err := s.valueOf(item)
		if err != nil {
			return nil, err
		}
		key := fmt.Sprintf("%v", v)
		groups[key] = append(listOf(groups[key]), item)
	}
	return groups, nil
}

// listOf returns v as a list, or an empty list when absent.
func listOf(v value.Value) []value.Value {
	if l, ok := v.([]value.Value); ok {
		return l
	}
	return nil
}

func (s reduceExec) aggregate(in pipeline.ResultSet) (value.Value, error) {
	nums, err := s.numbers(in)
	if err != nil {
		return nil, err
	}
	return aggregateNums(s.op, nums), nil
}

func (s reduceExec) numbers(in pipeline.ResultSet) ([]float64, error) {
	out := make([]float64, 0, len(in))
	for _, item := range in {
		v, err := s.valueOf(item)
		if err != nil {
			return nil, err
		}
		f, ferr := value.AsFloat(v)
		if ferr != nil {
			return nil, eval.ErrType
		}
		out = append(out, f)
	}
	return out, nil
}

// aggregateNums computes sum/avg/min/max over numbers; empty input yields null.
func aggregateNums(op ast.ReduceOp, nums []float64) value.Value {
	if len(nums) == 0 {
		return nil
	}
	switch op {
	case ast.OpSum:
		return sumOf(nums)
	case ast.OpAvg:
		return sumOf(nums) / float64(len(nums))
	case ast.OpMin:
		return slices.Min(nums)
	default:
		return slices.Max(nums)
	}
}

// sumOf totals a slice of floats.
func sumOf(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

// firstOrNil returns the first element or nil.
func firstOrNil(in pipeline.ResultSet) value.Value {
	if len(in) == 0 {
		return nil
	}
	return in[0]
}

// lastOrNil returns the last element or nil.
func lastOrNil(in pipeline.ResultSet) value.Value {
	if len(in) == 0 {
		return nil
	}
	return in[len(in)-1]
}

// sortExec reorders the result set by a key expression.
type sortExec struct {
	key  ast.Expr
	desc bool
	now  Clock
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
		return s.compare(a, b, &cmpErr)
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

// compare orders two keyed items, recording the first incomparable error.
func (s sortExec) compare(a, b keyedItem, cmpErr *error) int {
	c, err := value.Compare(a.key, b.key)
	if err != nil {
		*cmpErr = err
		return 0
	}
	if s.desc {
		return -c
	}
	return c
}

// unkey drops the sort keys, returning the ordered items.
func unkey(keyed []keyedItem) pipeline.ResultSet {
	out := make(pipeline.ResultSet, len(keyed))
	for i, k := range keyed {
		out[i] = k.item
	}
	return out
}

// limitExec truncates the result set to at most n elements.
type limitExec struct{ n int }

func (s limitExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	if s.n < len(in) {
		return in[:s.n], nil
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
