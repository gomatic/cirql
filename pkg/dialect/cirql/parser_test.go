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

func TestParse_Expressions(t *testing.T) {
	cases := []struct {
		want ast.Expr
		q    string
	}{
		{want: ast.BinaryExpr{Op: ast.OpEq}, q: `filter .a == 1`},
		{want: ast.BinaryExpr{Op: ast.OpNe}, q: `filter .a != 1`},
		{want: ast.BinaryExpr{Op: ast.OpGt}, q: `filter .a > 1`},
		{want: ast.BinaryExpr{Op: ast.OpLt}, q: `filter .a < 1`},
		{want: ast.BinaryExpr{Op: ast.OpGe}, q: `filter .a >= 1`},
		{want: ast.BinaryExpr{Op: ast.OpLe}, q: `filter .a <= 1`},
		{want: ast.BinaryExpr{Op: ast.OpAdd}, q: `filter .a + 1`},
		{want: ast.BinaryExpr{Op: ast.OpSub}, q: `filter .a - 1`},
		{want: ast.BinaryExpr{Op: ast.OpMul}, q: `filter .a * 1`},
		{want: ast.BinaryExpr{Op: ast.OpDiv}, q: `filter .a / 1`},
		{want: ast.BinaryExpr{Op: ast.OpMod}, q: `filter .a % 1`},
		{want: ast.BinaryExpr{Op: ast.OpAnd}, q: `filter .a && .b`},
		{want: ast.BinaryExpr{Op: ast.OpOr}, q: `filter .a || .b`},
		{want: ast.UnaryExpr{Op: ast.OpNot}, q: `filter !.a`},
		{want: ast.UnaryExpr{Op: ast.OpNeg}, q: `filter -.a`},
		{want: ast.FieldAccess{}, q: `filter (.a)`},
		{want: ast.VarRef{}, q: `filter $v`},
		{want: ast.FuncCall{}, q: `filter f(.a, .b)`},
	}
	for _, c := range cases {
		t.Run(c.q, func(t *testing.T) {
			got := c.q
			p := mustParse(t, got)
			cond := p.Stages[0].(ast.FilterStage).Cond
			assertExprShape(t, cond, c.want)
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

func TestParse_Literals(t *testing.T) {
	cases := map[string]any{
		`map { v: "hi" }`:   "hi",
		`map { v: 42 }`:     int64(42),
		`map { v: 3.5 }`:    3.5,
		`map { v: true }`:   true,
		`map { v: false }`:  false,
		`map { v: null }`:   nil,
		`map { v: "a\nb" }`: "a\nb",
		`map { v: "a\tb" }`: "a\tb",
		`map { v: "a\rb" }`: "a\rb",
		`map { v: "a\"b" }`: "a\"b",
		`map { v: "a\\b" }`: "a\\b",
		`map { v: "a\zb" }`: "azb",
	}
	for q, want := range cases {
		t.Run(q, func(t *testing.T) {
			p := mustParse(t, q)
			lit := p.Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.Literal)
			if lit.V != want {
				t.Fatalf("literal = %#v want %#v", lit.V, want)
			}
		})
	}
}

func TestParse_IntOverflowFallsBackToFloat(t *testing.T) {
	p := mustParse(t, `map { v: 99999999999999999999999999 }`)
	lit := p.Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.Literal)
	if _, ok := lit.V.(float64); !ok {
		t.Fatalf("overflowing int literal = %T want float64", lit.V)
	}
}

func TestParse_LimitOverflowKeepsAll(t *testing.T) {
	p := mustParse(t, `limit 99999999999999999999`)
	maxInt := int(^uint(0) >> 1)
	if n := p.Stages[0].(ast.LimitStage).N; n != maxInt {
		t.Fatalf("overflow limit N = %d want maxInt", n)
	}
}

func TestParse_FieldAccessPaths(t *testing.T) {
	p := mustParse(t, `map { v: .a.b[] }`)
	fa := p.Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.FieldAccess)
	if len(fa.Path) != 3 || fa.Path[0].Name != "a" || fa.Path[1].Name != "b" || !fa.Path[2].IsIter {
		t.Fatalf("path = %#v", fa.Path)
	}
	id := mustParse(t, `map { v: . }`).Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.FieldAccess)
	if len(id.Path) != 0 {
		t.Fatalf("identity path = %#v want empty", id.Path)
	}
}

func TestParse_Variable(t *testing.T) {
	p := mustParse(t, `filter $threshold > 1`)
	bin := p.Stages[0].(ast.FilterStage).Cond.(ast.BinaryExpr)
	if v, ok := bin.L.(ast.VarRef); !ok || v.Name != "threshold" {
		t.Fatalf("var = %#v", bin.L)
	}
}

func TestParse_FuncCallArgs(t *testing.T) {
	p := mustParse(t, `map { v: split(.tags, ",") }`)
	call := p.Stages[0].(ast.MapStage).Mappings[0].Expr.(ast.FuncCall)
	if call.Name != "split" || len(call.Args) != 2 {
		t.Fatalf("call = %#v", call)
	}
}

func TestParse_HTTPVariableURL(t *testing.T) {
	p := mustParse(t, `http $url`)
	h := p.Stages[0].(ast.HTTPStage)
	if _, ok := h.URL.(ast.VarRef); !ok {
		t.Fatalf("http url = %T want VarRef", h.URL)
	}
	p2 := mustParse(t, `http "http://x"`)
	if _, ok := p2.Stages[0].(ast.HTTPStage).URL.(ast.Literal); !ok {
		t.Fatal("http literal url not a Literal")
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
