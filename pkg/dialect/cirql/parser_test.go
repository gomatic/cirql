package cirql

import (
	"errors"
	"testing"

	"github.com/gomatic/cirql/ast"
)

func mustParse(t *testing.T, q string) ast.Pipeline {
	t.Helper()
	p, err := Parse(Query(q))
	if err != nil {
		t.Fatalf("Parse(%q): %v", q, err)
	}
	return p
}

func TestError_Message(t *testing.T) {
	if ErrParse.Error() != "cirql: parse error" {
		t.Fatalf("unexpected message %q", ErrParse.Error())
	}
}

func TestParse_TransformPipeline(t *testing.T) {
	p := mustParse(t, `filter .stars > 1000 | map { name: .name } | sort .name desc | limit 10`)
	if len(p.Stages) != 4 {
		t.Fatalf("got %d stages want 4", len(p.Stages))
	}
	if _, ok := p.Stages[0].(ast.FilterStage); !ok {
		t.Fatalf("stage0 = %T want FilterStage", p.Stages[0])
	}
	if _, ok := p.Stages[1].(ast.MapStage); !ok {
		t.Fatalf("stage1 = %T want MapStage", p.Stages[1])
	}
	if s := p.Stages[2].(ast.SortStage); !s.IsDesc {
		t.Fatal("sort should be desc")
	}
	if s := p.Stages[3].(ast.LimitStage); s.N != 10 {
		t.Fatalf("limit N = %d want 10", s.N)
	}
}

func TestParse_AllStageKinds(t *testing.T) {
	cases := []struct {
		want any
		q    string
	}{
		{want: ast.MapStage{}, q: `map { a: .a }`},
		{want: ast.FlatMapStage{}, q: `flatMap { a: .a }`},
		{want: ast.FilterStage{}, q: `filter .a`},
		{want: ast.ReduceStage{}, q: `reduce count`},
		{want: ast.ReduceStage{}, q: `reduce sum(.a)`},
		{want: ast.SortStage{}, q: `sort .a`},
		{want: ast.SortStage{}, q: `sort .a asc`},
		{want: ast.LimitStage{}, q: `limit 5`},
		{want: ast.UniqStage{}, q: `uniq`},
		{want: ast.UniqStage{}, q: `uniq .a`},
		{want: ast.StdinStage{}, q: `stdin`},
		{want: ast.FileStage{}, q: `file "x.json"`},
		{want: ast.HTTPStage{}, q: `http "http://h"`},
		{want: ast.HTTPStage{}, q: `http $u`},
		{want: ast.QueryStage{}, q: `query { a b }`},
		{want: ast.QueryStage{}, q: `{ a b }`},
	}
	for _, c := range cases {
		t.Run(c.q, func(t *testing.T) {
			p := mustParse(t, c.q)
			got := p.Stages[0]
			if gotT, wantT := typeName(got), typeName(c.want); gotT != wantT {
				t.Fatalf("stage = %s want %s", gotT, wantT)
			}
		})
	}
}

func typeName(v any) string {
	switch v.(type) {
	case ast.MapStage:
		return "map"
	case ast.FlatMapStage:
		return "flatmap"
	case ast.FilterStage:
		return "filter"
	case ast.ReduceStage:
		return "reduce"
	case ast.SortStage:
		return "sort"
	case ast.LimitStage:
		return "limit"
	case ast.UniqStage:
		return "uniq"
	case ast.StdinStage:
		return "stdin"
	case ast.FileStage:
		return "file"
	case ast.HTTPStage:
		return "http"
	case ast.QueryStage:
		return "query"
	}
	return "?"
}

func TestParse_ReduceOps(t *testing.T) {
	ops := map[string]ast.ReduceOp{
		"count": ast.OpCount, "sum(.a)": ast.OpSum, "min(.a)": ast.OpMin,
		"max(.a)": ast.OpMax, "avg(.a)": ast.OpAvg, "first": ast.OpFirst,
		"last": ast.OpLast, "group_by(.a)": ast.OpGroupBy, "collect(.a)": ast.OpCollect,
	}
	for src, want := range ops {
		t.Run(src, func(t *testing.T) {
			p := mustParse(t, "reduce "+src)
			if got := p.Stages[0].(ast.ReduceStage).Op; got != want {
				t.Fatalf("op = %q want %q", got, want)
			}
		})
	}
}

func assertExprShape(t *testing.T, got, want ast.Expr) {
	t.Helper()
	switch w := want.(type) {
	case ast.BinaryExpr:
		bin, ok := got.(ast.BinaryExpr)
		if !ok || bin.Op != w.Op {
			t.Fatalf("got %T(%v) want BinaryExpr(%v)", got, opOf(got), w.Op)
		}
	case ast.UnaryExpr:
		un, ok := got.(ast.UnaryExpr)
		if !ok || un.Op != w.Op {
			t.Fatalf("got %T want UnaryExpr(%v)", got, w.Op)
		}
	default:
		if typeOfExpr(got) != typeOfExpr(want) {
			t.Fatalf("got %T want %T", got, want)
		}
	}
}

func opOf(e ast.Expr) ast.BinOp {
	if b, ok := e.(ast.BinaryExpr); ok {
		return b.Op
	}
	return ""
}

func typeOfExpr(e ast.Expr) string {
	switch e.(type) {
	case ast.FieldAccess:
		return "field"
	case ast.VarRef:
		return "var"
	case ast.FuncCall:
		return "call"
	case ast.Literal:
		return "lit"
	}
	return "?"
}

func TestParse_IntOverflowFallsBackToFloat(t *testing.T) {
	p := mustParse(t, `map { v: 99999999999999999999999999 }`)
	lit := p.Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.Literal)
	if _, ok := lit.V.(float64); !ok {
		t.Fatalf("overflowing int literal = %T want float64", lit.V)
	}
}

func TestParse_FilePath(t *testing.T) {
	p := mustParse(t, `file "data.json"`)
	if p.Stages[0].(ast.FileStage).Path != "data.json" {
		t.Fatalf("path = %q", p.Stages[0].(ast.FileStage).Path)
	}
}

func TestParse_SyntaxError(t *testing.T) {
	for _, q := range []string{`map {`, `filter`, `| | |`, `sort`} {
		if _, err := Parse(Query(q)); !errors.Is(err, ErrParse) {
			t.Fatalf("Parse(%q) err = %v want ErrParse", q, err)
		}
	}
}

// A trailing sort direction keyword is NOT absorbed into the sort key path now
// that keywords are valid field names — .name is one segment, desc is the dir.
func TestParse_SortKeyDoesNotSwallowDirection(t *testing.T) {
	p := mustParse(t, `sort .name desc`)
	s := p.Stages[0].(ast.SortStage)
	fa := s.Key.(ast.FieldAccess)
	if len(fa.Path) != 1 || fa.Path[0].Name != "name" || !s.IsDesc {
		t.Fatalf("sort key=%#v desc=%v want .name desc", fa, s.IsDesc)
	}
}

// Deep field paths, root iteration, and mid-path iteration parse to the right
// segment lists.
func TestParse_PathShapes(t *testing.T) {
	cases := map[string][]ast.PathSegment{
		`filter .a.b.c`: {{Name: "a"}, {Name: "b"}, {Name: "c"}},
		`filter .[]`:    {{IsIter: true}},
		`filter .a[].b`: {{Name: "a"}, {IsIter: true}, {Name: "b"}},
		`filter .`:      nil,
	}
	for q, want := range cases {
		p := mustParse(t, q)
		fa := p.Stages[0].(ast.FilterStage).Cond.(ast.FieldAccess)
		if len(fa.Path) != len(want) {
			t.Fatalf("%s: path=%#v want %#v", q, fa.Path, want)
		}
		for i := range want {
			if fa.Path[i] != want[i] {
				t.Fatalf("%s: seg %d = %#v want %#v", q, i, fa.Path[i], want[i])
			}
		}
	}
}

// condOf parses `filter <expr>` and returns the filter condition.
func condOf(t *testing.T, q string) ast.Expr {
	t.Helper()
	return mustParse(t, q).Stages[0].(ast.FilterStage).Cond
}

// A syntax error carries the source position in its message (design contract:
// "ErrParse with line:col").
func TestParse_ErrorCarriesPosition(t *testing.T) {
	_, err := Parse(Query("filter @"))
	if !errors.Is(err, ErrParse) {
		t.Fatalf("got %v want ErrParse", err)
	}
	if got := err.Error(); !contains(got, "at 1:") {
		t.Errorf("error %q missing line:col position", got)
	}
}

// A query over the size bound is rejected before parsing, so adversarial deep
// nesting cannot overflow the stack.
func TestParse_QueryTooLarge(t *testing.T) {
	big := "filter " + repeat("(", MaxQueryBytes)
	_, err := Parse(Query(big))
	if !errors.Is(err, ErrQueryTooLarge) {
		t.Fatalf("got %v want ErrQueryTooLarge", err)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func repeat(s string, n int) string {
	b := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		b = append(b, s...)
	}
	return string(b)
}
