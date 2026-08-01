package stage

import (
	"errors"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
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
		{Key: "qty", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "qtys"}, {IsIter: true}}}},
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
		{Key: "x", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "xs"}, {IsIter: true}}}},
		{Key: "y", Expr: ast.FieldAccess{Path: []ast.PathSegment{{Name: "ys"}, {IsIter: true}}}},
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

// A nil condition in a programmatically built filter surfaces eval's ErrNilExpr
// as an error, not a panic.
func TestFilterNilCondErrors(t *testing.T) {
	st, err := Build(ast.FilterStage{Cond: nil}, nil)
	if err != nil {
		t.Fatalf("unexpected build error: %v", err)
	}
	if _, err := st.Execute(pipeline.ResultSet{map[string]value.Value{}}); !errors.Is(err, eval.ErrNilExpr) {
		t.Errorf("got %v, want eval.ErrNilExpr", err)
	}
}

// flatMap over an empty list produces ZERO rows, not a phantom {key: null} row
// (spec §5.3: one output object per list element).
func TestFlatMap_EmptyListYieldsNoRows(t *testing.T) {
	in := pipeline.ResultSet{obj("xs", []value.Value{})}
	out := run(t, ast.FlatMapStage{Mappings: []ast.Mapping{{Key: "x", Expr: field("xs")}}}, in)
	if len(out) != 0 {
		t.Fatalf("got %d rows, want 0", len(out))
	}
}

// flatMap with only scalar mappings still emits exactly one row per input.
func TestFlatMap_AllScalarYieldsOneRow(t *testing.T) {
	in := pipeline.ResultSet{obj("a", int64(1))}
	out := run(t, ast.FlatMapStage{Mappings: []ast.Mapping{{Key: "x", Expr: field("a")}}}, in)
	if len(out) != 1 {
		t.Fatalf("got %d rows, want 1", len(out))
	}
}

// flatMap expands one row per list element, scalars repeating.
func TestFlatMap_ExpandsPerElement(t *testing.T) {
	in := pipeline.ResultSet{obj("xs", []value.Value{int64(1), int64(2), int64(3)}, "k", "c")}
	out := run(t, ast.FlatMapStage{Mappings: []ast.Mapping{
		{Key: "x", Expr: field("xs")},
		{Key: "k", Expr: field("k")},
	}}, in)
	if len(out) != 3 {
		t.Fatalf("got %d rows, want 3", len(out))
	}
	for i, want := range []int64{1, 2, 3} {
		row := out[i].(map[string]value.Value)
		if row["x"] != want || row["k"] != "c" {
			t.Errorf("row %d = %#v, want x=%d k=c", i, row, want)
		}
	}
}

func keysOf(m map[string]value.Value) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

// filter keeps the RIGHT items, not merely the right count.
func TestFilter_KeepsMatchingItems(t *testing.T) {
	in := pipeline.ResultSet{obj("n", int64(0)), obj("n", int64(5)), obj("n", int64(9))}
	cond := ast.BinaryExpr{Op: ast.OpGt, L: field("n"), R: ast.Literal{V: int64(1)}}
	out := run(t, ast.FilterStage{Cond: cond}, in)
	if len(out) != 2 {
		t.Fatalf("got %d, want 2", len(out))
	}
	if out[0].(map[string]value.Value)["n"] != int64(5) || out[1].(map[string]value.Value)["n"] != int64(9) {
		t.Errorf("kept the wrong items: %#v", out)
	}
}
