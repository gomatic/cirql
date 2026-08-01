package stage

import (
	"encoding/json"
	"slices"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
)

// The reduce stage: count, sum, min, max, avg, first, last, group_by and
// collect. It is the one stage whose output shape depends on the operator
// rather than on the input, so it is the largest and is kept apart from the
// element-wise stages that surround it.

// reduceExec aggregates the result set into a single-element set.
type reduceExec struct {
	arg ast.Expr
	now Clock
	op  ast.ReduceOp
}

func (s reduceExec) Execute(in pipeline.ResultSet) (pipeline.ResultSet, error) {
	v, err := s.reduce(in)
	if err != nil {
		return nil, err
	}
	return pipeline.ResultSet{v}, nil
}

// reduce dispatches on the operator. Every declared ReduceOp is listed, so
// adding one to the AST without deciding what it reduces to is a visible
// omission rather than a silent fall-through: the default previously sent every
// unrecognised operator to the numeric aggregate, which answered a numeric
// question for an operator that had never been given a meaning.
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
	case ast.OpSum, ast.OpAvg, ast.OpMin, ast.OpMax:
		return s.aggregate(in)
	default:
		return nil, eval.ErrType
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
		key, kerr := groupKey(v)
		if kerr != nil {
			return nil, kerr
		}
		groups[key] = append(listOf(groups[key]), item)
	}
	return groups, nil
}

// groupKey renders a group-by key value as a JSON object key: a string value
// keys directly (so group_by over a name field yields {"alice": …}); every
// other value is rendered as its JSON text, so objects and booleans never leak
// Go syntax (fmt %v would render an object as "map[...]"). encoding/json emits
// map keys in sorted order, so an object key is deterministic.
func groupKey(v value.Value) (string, error) {
	if s, ok := v.(string); ok {
		return s, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", eval.ErrType
	}
	return string(b), nil
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
	return aggregateNums(aggregation(s.op), nums), nil
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

// aggregation is the subset of ast.ReduceOp that folds a list of numbers into
// one. The narrow type is what lets aggregateNums be total over its own domain:
// reduce dispatches the other five operators elsewhere, so a switch over the
// whole of ast.ReduceOp here would have to carry five arms nothing can reach.
type aggregation ast.ReduceOp

const (
	aggSum aggregation = aggregation(ast.OpSum)
	aggAvg aggregation = aggregation(ast.OpAvg)
	aggMin aggregation = aggregation(ast.OpMin)
	aggMax aggregation = aggregation(ast.OpMax)
)

// aggregateNums computes sum/avg/min/max over numbers; empty input yields null.
// An aggregation outside the declared four has no numeric meaning and yields
// null rather than silently answering with the maximum, which is what a default
// arm did before the operators were named.
func aggregateNums(op aggregation, nums []float64) value.Value {
	if len(nums) == 0 {
		return nil
	}
	switch op {
	case aggSum:
		return sumOf(nums)
	case aggAvg:
		return sumOf(nums) / float64(len(nums))
	case aggMin:
		return slices.Min(nums)
	case aggMax:
		return slices.Max(nums)
	default:
		return nil
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
