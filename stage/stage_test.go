package stage

import (
	"errors"
	"testing"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
	"github.com/gomatic/cirql/value"
)

func obj(kv ...value.Value) map[string]value.Value {
	m := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		m[kv[i].(string)] = kv[i+1]
	}
	return m
}

func field(name string) ast.Expr {
	return ast.FieldAccess{Path: []ast.PathSegment{{Name: name}}}
}

func build(t *testing.T, s ast.Stage) pipeline.Stage {
	t.Helper()
	st, err := Build(s, nil)
	if err != nil {
		t.Fatalf("Build(%T): %v", s, err)
	}
	return st
}

func run(t *testing.T, s ast.Stage, in pipeline.ResultSet) pipeline.ResultSet {
	t.Helper()
	out, err := build(t, s).Execute(in)
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
	return out
}

func TestError_Message(t *testing.T) {
	if ErrStageUnsupported.Error() != "cirql: source stage not supported in core" {
		t.Fatalf("unexpected message %q", ErrStageUnsupported.Error())
	}
}

func TestBuild_SourceStagesUnsupported(t *testing.T) {
	for _, s := range []ast.Stage{ast.FileStage{}, ast.HTTPStage{}, ast.QueryStage{}} {
		if _, err := Build(s, nil); !errors.Is(err, ErrStageUnsupported) {
			t.Fatalf("Build(%T) err=%v want ErrStageUnsupported", s, err)
		}
	}
}

func TestStdin_Identity(t *testing.T) {
	in := pipeline.ResultSet{obj("a", int64(1))}
	out := run(t, ast.StdinStage{}, in)
	if len(out) != 1 {
		t.Fatalf("stdin len=%d want 1", len(out))
	}
}

func TestMap(t *testing.T) {
	in := pipeline.ResultSet{obj("name", "x", "extra", int64(9))}
	out := run(t, ast.MapStage{Mappings: []ast.Mapping{{Key: "n", Expr: field("name")}}}, in)
	got := out[0].(map[string]value.Value)
	if got["n"] != "x" || len(got) != 1 {
		t.Fatalf("map=%#v", got)
	}
}

func TestMap_PropagatesError(t *testing.T) {
	in := pipeline.ResultSet{obj("a", int64(1))}
	_, err := build(t, ast.MapStage{Mappings: []ast.Mapping{
		{Key: "n", Expr: ast.FuncCall{Name: "nope"}},
	}}).Execute(in)
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestFilter(t *testing.T) {
	in := pipeline.ResultSet{
		obj("n", int64(0)),
		obj("n", int64(5)),
	}
	cond := ast.BinaryExpr{Op: ast.OpGt, L: field("n"), R: ast.Literal{V: int64(1)}}
	out := run(t, ast.FilterStage{Cond: cond}, in)
	if len(out) != 1 {
		t.Fatalf("filter len=%d want 1", len(out))
	}
}

func TestFilter_PropagatesError(t *testing.T) {
	in := pipeline.ResultSet{obj("a", int64(1))}
	_, err := build(t, ast.FilterStage{Cond: ast.FuncCall{Name: "nope"}}).Execute(in)
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestFlatMap_Expands(t *testing.T) {
	item := obj("title", "p", "qtys", []value.Value{int64(1), int64(2), int64(3)})
	stage := ast.FlatMapStage{Mappings: []ast.Mapping{
		{Key: "title", Expr: field("title")},
		{Key: "qty", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "qtys"}, {Iter: true}}}},
	}}
	out := run(t, stage, pipeline.ResultSet{item})
	if len(out) != 3 {
		t.Fatalf("flatMap len=%d want 3", len(out))
	}
	first := out[0].(map[string]value.Value)
	if first["title"] != "p" || first["qty"] != int64(1) {
		t.Fatalf("row0=%#v", first)
	}
}

func TestFlatMap_ScalarOnlyActsLikeMap(t *testing.T) {
	out := run(t, ast.FlatMapStage{Mappings: []ast.Mapping{{Key: "n", Expr: field("a")}}},
		pipeline.ResultSet{obj("a", int64(1))})
	if len(out) != 1 {
		t.Fatalf("len=%d want 1", len(out))
	}
}

func TestFlatMap_ShorterListPadsNil(t *testing.T) {
	item := obj("xs", []value.Value{int64(1)}, "ys", []value.Value{int64(7), int64(8)})
	stage := ast.FlatMapStage{Mappings: []ast.Mapping{
		{Key: "x", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "xs"}, {Iter: true}}}},
		{Key: "y", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "ys"}, {Iter: true}}}},
	}}
	out := run(t, stage, pipeline.ResultSet{item})
	if len(out) != 2 {
		t.Fatalf("len=%d want 2", len(out))
	}
	if out[1].(map[string]value.Value)["x"] != nil {
		t.Fatalf("row1.x = %v want nil", out[1].(map[string]value.Value)["x"])
	}
}

func TestFlatMap_PropagatesError(t *testing.T) {
	_, err := build(t, ast.FlatMapStage{Mappings: []ast.Mapping{
		{Key: "n", Expr: ast.FuncCall{Name: "nope"}},
	}}).Execute(pipeline.ResultSet{obj("a", int64(1))})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestReduce_Count(t *testing.T) {
	out := run(t, ast.ReduceStage{Op: ast.OpCount}, pipeline.ResultSet{obj(), obj(), obj()})
	if out[0] != int64(3) {
		t.Fatalf("count=%v want 3", out[0])
	}
}

func TestReduce_FirstLast(t *testing.T) {
	in := pipeline.ResultSet{obj("i", int64(1)), obj("i", int64(2))}
	if run(t, ast.ReduceStage{Op: ast.OpFirst}, in)[0].(map[string]value.Value)["i"] != int64(1) {
		t.Fatal("first wrong")
	}
	if run(t, ast.ReduceStage{Op: ast.OpLast}, in)[0].(map[string]value.Value)["i"] != int64(2) {
		t.Fatal("last wrong")
	}
	if run(t, ast.ReduceStage{Op: ast.OpFirst}, pipeline.ResultSet{})[0] != nil {
		t.Fatal("first empty should be nil")
	}
	if run(t, ast.ReduceStage{Op: ast.OpLast}, pipeline.ResultSet{})[0] != nil {
		t.Fatal("last empty should be nil")
	}
}

func TestReduce_SumAvgMinMax(t *testing.T) {
	in := pipeline.ResultSet{obj("v", int64(2)), obj("v", int64(4))}
	cases := map[ast.ReduceOp]value.Value{
		ast.OpSum: 6.0, ast.OpAvg: 3.0, ast.OpMin: 2.0, ast.OpMax: 4.0,
	}
	for op, want := range cases {
		out := run(t, ast.ReduceStage{Op: op, Arg: field("v")}, in)
		if out[0] != want {
			t.Fatalf("%s=%v want %v", op, out[0], want)
		}
	}
}

func TestReduce_AggregateEmptyIsNull(t *testing.T) {
	out := run(t, ast.ReduceStage{Op: ast.OpSum, Arg: field("v")}, pipeline.ResultSet{})
	if out[0] != nil {
		t.Fatalf("sum empty=%v want nil", out[0])
	}
}

func TestReduce_NoArgUsesItem(t *testing.T) {
	out := run(t, ast.ReduceStage{Op: ast.OpSum}, pipeline.ResultSet{int64(1), int64(2)})
	if out[0] != 3.0 {
		t.Fatalf("sum no-arg=%v want 3", out[0])
	}
}

func TestReduce_AggregateTypeError(t *testing.T) {
	_, err := build(t, ast.ReduceStage{Op: ast.OpSum, Arg: field("v")}).
		Execute(pipeline.ResultSet{obj("v", "x")})
	if !errors.Is(err, eval.ErrType) {
		t.Fatalf("err=%v want ErrType", err)
	}
}

func TestReduce_AggregatePropagatesEvalError(t *testing.T) {
	_, err := build(t, ast.ReduceStage{Op: ast.OpSum, Arg: ast.FuncCall{Name: "nope"}}).
		Execute(pipeline.ResultSet{obj()})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestReduce_Collect(t *testing.T) {
	in := pipeline.ResultSet{obj("v", int64(1)), obj("v", int64(2))}
	out := run(t, ast.ReduceStage{Op: ast.OpCollect, Arg: field("v")}, in)
	list := out[0].([]value.Value)
	if len(list) != 2 || list[0] != int64(1) {
		t.Fatalf("collect=%#v", list)
	}
}

func TestReduce_CollectPropagatesError(t *testing.T) {
	_, err := build(t, ast.ReduceStage{Op: ast.OpCollect, Arg: ast.FuncCall{Name: "nope"}}).
		Execute(pipeline.ResultSet{obj()})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestReduce_GroupBy(t *testing.T) {
	in := pipeline.ResultSet{
		obj("k", "a", "n", int64(1)),
		obj("k", "b", "n", int64(2)),
		obj("k", "a", "n", int64(3)),
	}
	out := run(t, ast.ReduceStage{Op: ast.OpGroupBy, Arg: field("k")}, in)
	groups := out[0].(map[string]value.Value)
	if len(groups["a"].([]value.Value)) != 2 || len(groups["b"].([]value.Value)) != 1 {
		t.Fatalf("groupBy=%#v", groups)
	}
}

func TestReduce_GroupByPropagatesError(t *testing.T) {
	_, err := build(t, ast.ReduceStage{Op: ast.OpGroupBy, Arg: ast.FuncCall{Name: "nope"}}).
		Execute(pipeline.ResultSet{obj()})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestSort_AscDesc(t *testing.T) {
	in := pipeline.ResultSet{obj("n", int64(2)), obj("n", int64(1)), obj("n", int64(3))}
	asc := run(t, ast.SortStage{Key: field("n")}, in)
	if asc[0].(map[string]value.Value)["n"] != int64(1) {
		t.Fatalf("asc first=%v want 1", asc[0])
	}
	desc := run(t, ast.SortStage{Key: field("n"), Desc: true}, in)
	if desc[0].(map[string]value.Value)["n"] != int64(3) {
		t.Fatalf("desc first=%v want 3", desc[0])
	}
}

func TestSort_KeyEvalError(t *testing.T) {
	_, err := build(t, ast.SortStage{Key: ast.FuncCall{Name: "nope"}}).
		Execute(pipeline.ResultSet{obj(), obj()})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}

func TestSort_IncomparableError(t *testing.T) {
	in := pipeline.ResultSet{obj("n", "x"), obj("n", int64(1))}
	_, err := build(t, ast.SortStage{Key: field("n")}).Execute(in)
	if !errors.Is(err, value.ErrIncomparable) {
		t.Fatalf("err=%v want ErrIncomparable", err)
	}
}

func TestLimit(t *testing.T) {
	in := pipeline.ResultSet{obj(), obj(), obj()}
	if len(run(t, ast.LimitStage{N: 2}, in)) != 2 {
		t.Fatal("limit 2 failed")
	}
	if len(run(t, ast.LimitStage{N: 10}, in)) != 3 {
		t.Fatal("limit beyond len should keep all")
	}
}

func TestUniq_WholeValue(t *testing.T) {
	in := pipeline.ResultSet{int64(1), int64(1), int64(2)}
	out := run(t, ast.UniqStage{}, in)
	if len(out) != 2 {
		t.Fatalf("uniq len=%d want 2", len(out))
	}
}

func TestUniq_ByKey(t *testing.T) {
	in := pipeline.ResultSet{
		obj("k", "a"), obj("k", "a"), obj("k", "b"),
	}
	out := run(t, ast.UniqStage{Key: field("k")}, in)
	if len(out) != 2 {
		t.Fatalf("uniq-by-key len=%d want 2", len(out))
	}
}

func TestUniq_KeyEvalError(t *testing.T) {
	_, err := build(t, ast.UniqStage{Key: ast.FuncCall{Name: "nope"}}).
		Execute(pipeline.ResultSet{obj()})
	if !errors.Is(err, eval.ErrUnknownFunc) {
		t.Fatalf("err=%v want ErrUnknownFunc", err)
	}
}
