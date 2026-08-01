// Package stage implements the cirql transform stages and Build, which maps a
// parsed stage node to its executor. Source stages (file/http/query) report
// ErrStageUnsupported until the sources sub-project provides them; the stdin
// source is the identity stage (its input is the pipeline's input).
package stage

import (
	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
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
		return sortExec{key: n.Key, isDesc: n.IsDesc, now: now}, nil
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
	now      Clock
	mappings []ast.Mapping
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
	now      Clock
	mappings []ast.Mapping
}

// kv is one evaluated mapping.
type kv struct {
	val value.Value
	key string
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
		rows[i] = rowAt(ev, rowIndex(i))
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

// fanout is the number of output rows. With no list-valued mapping the stage is
// a plain map (one row); with any list mapping it is the longest list's length,
// so an empty list produces ZERO rows (spec §5.3: one row per element) rather
// than a phantom null row.
func fanout(ev []kv) int {
	hasList := false
	n := 0
	for _, e := range ev {
		if l, ok := e.val.([]value.Value); ok {
			hasList = true
			if len(l) > n {
				n = len(l)
			}
		}
	}
	if !hasList {
		return 1
	}
	return n
}

// rowIndex is the position of an expanded output row within a flatMap fanout.
type rowIndex int

// rowAt builds the i-th expanded row: list mappings take element i, scalars repeat.
func rowAt(ev []kv, i rowIndex) value.Value {
	row := make(map[string]value.Value, len(ev))
	for _, e := range ev {
		row[e.key] = elemAt(e.val, i)
	}
	return row
}

// elemAt returns the i-th element of a list (nil past the end) or a scalar as-is.
func elemAt(v value.Value, i rowIndex) value.Value {
	l, ok := v.([]value.Value)
	if !ok {
		return v
	}
	if int(i) < len(l) {
		return l[int(i)]
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
