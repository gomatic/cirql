package stage

import (
	"errors"
	"math"
	"testing"

	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	"github.com/gomatic/cirql/eval"
	"github.com/gomatic/cirql/pipeline"
)

func TestSort_AscDesc(t *testing.T) {
	in := pipeline.ResultSet{obj("n", int64(2)), obj("n", int64(1)), obj("n", int64(3))}
	asc := run(t, ast.SortStage{Key: field("n")}, in)
	if asc[0].(map[string]value.Value)["n"] != int64(1) {
		t.Fatalf("asc first=%v want 1", asc[0])
	}
	desc := run(t, ast.SortStage{Key: field("n"), IsDesc: true}, in)
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

// A negative limit — constructible only programmatically, never by the parser —
// behaves as limit 0 instead of panicking.
func TestLimitNegativeYieldsEmpty(t *testing.T) {
	st, err := Build(ast.LimitStage{N: -1}, nil)
	if err != nil {
		t.Fatalf("unexpected build error: %v", err)
	}
	out, err := st.Execute(pipeline.ResultSet{map[string]value.Value{"a": int64(1)}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) != 0 {
		t.Errorf("got %v, want empty", out)
	}
}

// sort is stable: equal keys keep input order (asc) and are NOT reversed by
// desc (desc negates the comparator, it does not reverse equal runs).
func TestSort_IsStable(t *testing.T) {
	in := pipeline.ResultSet{
		obj("k", int64(1), "id", "a"),
		obj("k", int64(1), "id", "b"),
		obj("k", int64(1), "id", "c"),
	}
	for _, desc := range []bool{false, true} {
		out := run(t, ast.SortStage{Key: field("k"), IsDesc: desc}, in)
		got := []string{
			out[0].(map[string]value.Value)["id"].(string),
			out[1].(map[string]value.Value)["id"].(string),
			out[2].(map[string]value.Value)["id"].(string),
		}
		if got[0] != "a" || got[1] != "b" || got[2] != "c" {
			t.Errorf("desc=%v stable order = %v, want [a b c]", desc, got)
		}
	}
}

// Whole-object uniq (no key) dedupes equal OBJECTS without panicking — the
// exact case that panicked under the old go-json Equal.
func TestUniq_WholeObjectDedupes(t *testing.T) {
	in := pipeline.ResultSet{
		obj("a", int64(1), "b", int64(2)),
		obj("a", int64(1), "b", int64(2)),
		obj("a", int64(9)),
	}
	out := run(t, ast.UniqStage{}, in)
	if len(out) != 2 {
		t.Fatalf("got %d, want 2 (dedupe identical objects)", len(out))
	}
}

// TestLimitExecTreatsANegativeCountAsZero names limitExec's claim. A negative n
// cannot come from the parser, but it can from a programmatically assembled
// pipeline — and `in[:n]` with a negative n panics. "Behaves as zero" is what
// makes an embedder's mistake an empty result rather than a crash inside the
// library.
func TestLimitExecTreatsANegativeCountAsZero(t *testing.T) {
	in := pipeline.ResultSet{
		map[string]value.Value{"a": int64(1)},
		map[string]value.Value{"a": int64(2)},
	}

	for _, n := range []int{-1, -100, math.MinInt} {
		out, err := limitExec{n: n}.Execute(in)
		if err != nil {
			t.Fatalf("limitExec{%d}: %v", n, err)
		}
		if len(out) != 0 {
			t.Fatalf("limitExec{%d} kept %d elements, want none", n, len(out))
		}
	}

	out, err := limitExec{n: 1}.Execute(in)
	if err != nil || len(out) != 1 {
		t.Fatalf("limitExec{1} kept %d (err %v), want 1", len(out), err)
	}
}

// TestUniqExecDeduplicatesByKeyWhenSetElseWholeValue names uniqExec's claim.
// The two modes answer different questions — "one row per user" versus "drop
// identical rows" — and confusing them silently discards data the query asked
// to keep. Order must survive too: uniq is not sort, so the first occurrence of
// each distinct value stays where it was.
func TestUniqExecDeduplicatesByKeyWhenSetElseWholeValue(t *testing.T) {
	in := pipeline.ResultSet{
		map[string]value.Value{"user": "a", "n": int64(1)},
		map[string]value.Value{"user": "b", "n": int64(2)},
		map[string]value.Value{"user": "a", "n": int64(3)},
		map[string]value.Value{"user": "b", "n": int64(2)},
	}

	whole, err := uniqExec{}.Execute(in)
	if err != nil {
		t.Fatal(err)
	}
	if len(whole) != 3 {
		t.Fatalf("whole-value uniq kept %d, want 3 — only the exact duplicate drops", len(whole))
	}
	if whole[0].(map[string]value.Value)["n"] != value.Value(int64(1)) {
		t.Fatal("uniq must preserve the order of first occurrence")
	}

	byUser, err := uniqExec{key: ast.FieldAccess{Path: []ast.PathSegment{{Name: "user"}}}}.Execute(in)
	if err != nil {
		t.Fatal(err)
	}
	if len(byUser) != 2 {
		t.Fatalf("uniq by key kept %d, want one per distinct user", len(byUser))
	}
	if byUser[0].(map[string]value.Value)["n"] != value.Value(int64(1)) {
		t.Fatal("uniq by key keeps the FIRST object per key, not the last")
	}
}
