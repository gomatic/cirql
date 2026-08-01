package stage

import (
	"errors"
	"math"
	"strings"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
)

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
	// A slice rather than a map keyed by ast.ReduceOp: this table covers the
	// four numeric aggregations deliberately, and an enum-keyed map would be
	// read as a claim to cover every reduce operator.
	cases := []struct {
		want value.Value
		op   ast.ReduceOp
	}{
		{op: ast.OpSum, want: 6.0},
		{op: ast.OpAvg, want: 3.0},
		{op: ast.OpMin, want: 2.0},
		{op: ast.OpMax, want: 4.0},
	}
	for _, tc := range cases {
		op, want := tc.op, tc.want
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

// group_by keys are JSON: a string value keys directly; a non-string value is
// rendered as JSON text (never Go's map[...] syntax) and distinct types do not
// silently collapse.
func TestReduce_GroupByKeysAreJSON(t *testing.T) {
	in := pipeline.ResultSet{
		obj("g", int64(1)),
		obj("g", "1"),
		obj("g", true),
	}
	out := run(t, ast.ReduceStage{Op: ast.OpGroupBy, Arg: field("g")}, in)
	groups := out[0].(map[string]value.Value)
	// int 1 -> "1", string "1" -> "1": these coincide as JSON object keys (an
	// inherent consequence of object-keyed grouping), so they share one group;
	// bool true -> "true" is its own group and never renders as a Go-ism.
	if _, ok := groups["true"]; !ok {
		t.Errorf("missing JSON key \"true\"; groups=%v keys", keysOf(groups))
	}
	for k := range groups {
		if k == "<nil>" || len(k) >= 4 && k[:4] == "map[" {
			t.Errorf("group key %q leaks Go syntax", k)
		}
	}
}

// A string group key is the bare string (no surrounding quotes) so results read
// naturally: group_by .name -> {"alice": [...]}.
func TestReduce_GroupByStringKeyIsBare(t *testing.T) {
	in := pipeline.ResultSet{obj("name", "alice"), obj("name", "alice"), obj("name", "bob")}
	out := run(t, ast.ReduceStage{Op: ast.OpGroupBy, Arg: field("name")}, in)
	groups := out[0].(map[string]value.Value)
	if len(groups["alice"].([]value.Value)) != 2 || len(groups["bob"].([]value.Value)) != 1 {
		t.Fatalf("groups=%#v", groups)
	}
}

// group_by over a key that is not JSON-serializable (a non-finite float,
// reachable via overflow arithmetic) surfaces ErrType rather than a bad key.
func TestReduce_GroupByNonSerializableKeyErrors(t *testing.T) {
	in := pipeline.ResultSet{obj("g", math.Inf(1))}
	_, err := build(t, ast.ReduceStage{Op: ast.OpGroupBy, Arg: field("g")}).Execute(in)
	if !errors.Is(err, eval.ErrType) {
		t.Fatalf("got %v, want eval.ErrType", err)
	}
}

// TestGroupKeyRendersNonStringsAsJSONNotGoSyntax names groupKey's claim. The
// key becomes a field name in the emitted object, so it is user-visible output,
// not an internal token. fmt's %v would render an object as `map[a:1]` — Go
// syntax leaking into a JSON document, unparseable by whatever consumes it —
// and a bool as `true` without quotes being distinguishable from the string
// "true". JSON text is the only rendering that round-trips.
func TestGroupKeyRendersNonStringsAsJSONNotGoSyntax(t *testing.T) {
	for _, tc := range []struct {
		in   value.Value
		name string
		want string
	}{
		{name: "a string keys directly", in: "alice", want: "alice"},
		{name: "an integer", in: int64(3), want: "3"},
		{name: "a bool", in: true, want: "true"},
		{name: "null", in: nil, want: "null"},
		{name: "a list", in: []value.Value{int64(1), "a"}, want: `[1,"a"]`},
		{name: "an object", in: map[string]value.Value{"a": int64(1)}, want: `{"a":1}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := groupKey(tc.in)
			if err != nil {
				t.Fatalf("groupKey(%#v): %v", tc.in, err)
			}
			if got != tc.want {
				t.Fatalf("groupKey(%#v) = %q, want %q", tc.in, got, tc.want)
			}
			if strings.Contains(got, "map[") {
				t.Fatalf("groupKey leaked Go syntax: %q", got)
			}
		})
	}
}

// TestGroupKeyIsDeterministicForObjectKeys names the second half of groupKey's
// claim: "encoding/json emits map keys in sorted order, so an object key is
// deterministic." Without that, two runs over identical data would emit
// different group names — the output would not be reproducible, and a
// downstream diff would show spurious changes.
func TestGroupKeyIsDeterministicForObjectKeys(t *testing.T) {
	object := map[string]value.Value{"z": int64(1), "a": int64(2), "m": int64(3)}

	first, err := groupKey(object)
	if err != nil {
		t.Fatal(err)
	}
	if first != `{"a":2,"m":3,"z":1}` {
		t.Fatalf("got %q, want keys in sorted order", first)
	}
	for i := range 50 {
		again, err := groupKey(map[string]value.Value{"z": int64(1), "a": int64(2), "m": int64(3)})
		if err != nil || again != first {
			t.Fatalf("iteration %d rendered %q (err %v), want the stable %q", i, again, err, first)
		}
	}
}

// TestReduceRejectsAnOperatorItWasNeverGivenAMeaning pins the dispatch's
// totality. The switch previously ended in a default that ran the NUMERIC
// aggregate, so a ReduceStage carrying an operator the AST had grown but this
// file had not — or one assembled programmatically — silently produced a number
// where the caller expected whatever that operator was supposed to mean. An
// operator with no defined reduction is a type error, not a sum.
func TestReduceRejectsAnOperatorItWasNeverGivenAMeaning(t *testing.T) {
	// The arg matters: without it the aggregate path fails on its own, for a
	// type error inside numbers(), and the test would pass under the very
	// fall-through it exists to forbid. With a numeric arg the aggregate path
	// SUCCEEDS, so the only way to get an error here is the refusal.
	in := pipeline.ResultSet{obj("v", int64(2)), obj("v", int64(4))}
	arg := field("v")

	sum, err := reduceExec{op: ast.OpSum, arg: arg}.Execute(in)
	if err != nil || sum[0] != 6.0 {
		t.Fatalf("the fixture must aggregate cleanly (got %v, err %v), or the refusal below proves nothing", sum, err)
	}

	_, err = reduceExec{op: ast.ReduceOp("undeclared"), arg: arg}.Execute(in)

	if !errors.Is(err, eval.ErrType) {
		t.Fatalf("got %v, want eval.ErrType — an undeclared operator must not silently aggregate", err)
	}
}

// TestAggregateNumsRejectsAnUndeclaredAggregation covers the same boundary one
// level down, where the old default arm silently answered every unrecognised
// aggregation with the MAXIMUM.
func TestAggregateNumsRejectsAnUndeclaredAggregation(t *testing.T) {
	if got := aggregateNums(aggregation("undeclared"), []float64{1, 5, 3}); got != nil {
		t.Fatalf("aggregateNums(undeclared) = %v, want null rather than a silent max", got)
	}
	if got := aggregateNums(aggMax, []float64{1, 5, 3}); got != 5.0 {
		t.Fatalf("aggregateNums(aggMax) = %v, want 5 — the declared operators still work", got)
	}
}
