package cirql

import (
	"testing"

	"github.com/gomatic/cirql/ast"
)

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

// Keywords are addressable as field names in path and mapping positions — a
// JSON query language must handle fields named count, type, map, first, etc.
func TestParse_KeywordsAsFieldNames(t *testing.T) {
	p := mustParse(t, `filter .count > 1 | map { first: .last, sum: .map }`)
	fs, ok := p.Stages[0].(ast.FilterStage)
	if !ok {
		t.Fatalf("stage0 = %T", p.Stages[0])
	}
	cmp, ok := fs.Cond.(ast.BinaryExpr)
	if !ok {
		t.Fatalf("cond = %T want BinaryExpr", fs.Cond)
	}
	fa, ok := cmp.L.(ast.FieldAccess)
	if !ok || len(fa.Path) != 1 || fa.Path[0].Name != "count" {
		t.Fatalf("lhs path = %#v want .count", fa)
	}
	ms := p.Stages[1].(ast.MapStage)
	if ms.Mappings[0].Key != "first" || ms.Mappings[1].Key != "sum" {
		t.Fatalf("mapping keys = %q,%q want first,sum", ms.Mappings[0].Key, ms.Mappings[1].Key)
	}
	// The mapping VALUE .map addresses a field named map, not a stage.
	fa2 := ms.Mappings[1].Expr.(ast.FieldAccess)
	if len(fa2.Path) != 1 || fa2.Path[0].Name != "map" {
		t.Fatalf("mapping[1] expr = %#v want .map", fa2)
	}
}

// Operator precedence and associativity are a core deliverable; pin the ladder
// (unary > mul > add > cmp > and > or) and left-associativity by asserting the
// parsed tree SHAPE, not just the top node.
func TestParse_PrecedenceAndAssociativity(t *testing.T) {
	// 1 + 2 * 3  ==>  1 + (2 * 3): top is Add, right is Mul.
	add := condOf(t, `filter 1 + 2 * 3`).(ast.BinaryExpr)
	if add.Op != ast.OpAdd {
		t.Fatalf("top op = %s want +", add.Op)
	}
	if r, ok := add.R.(ast.BinaryExpr); !ok || r.Op != ast.OpMul {
		t.Fatalf("rhs = %#v want a Mul", add.R)
	}
	// 8 - 4 - 2  ==>  (8 - 4) - 2: left-associative, so left is the nested Sub.
	sub := condOf(t, `filter 8 - 4 - 2`).(ast.BinaryExpr)
	if l, ok := sub.L.(ast.BinaryExpr); !ok || l.Op != ast.OpSub {
		t.Fatalf("lhs = %#v want a nested Sub (left-assoc)", sub.L)
	}
	// a > 1 && b < 2  ==>  (a>1) && (b<2): && binds looser than comparison.
	and := condOf(t, `filter .a > 1 && .b < 2`).(ast.BinaryExpr)
	if and.Op != ast.OpAnd {
		t.Fatalf("top op = %s want &&", and.Op)
	}
	// -2 * 3  ==>  (-2) * 3: unary binds tighter than *.
	mul := condOf(t, `filter -2 * 3`).(ast.BinaryExpr)
	if mul.Op != ast.OpMul {
		t.Fatalf("top op = %s want * (unary binds tighter)", mul.Op)
	}
	if _, ok := mul.L.(ast.UnaryExpr); !ok {
		t.Fatalf("lhs = %#v want a UnaryExpr", mul.L)
	}
}

// TestAtoiClampsAnOverflowingLimitToKeepAll names atoi's claim. The lexer
// guarantees digits, so the only way strconv fails here is overflow — and the
// answer must be "keep everything", not zero. Returning zero on overflow would
// turn `limit 99999999999999999999` into `limit 0` and silently produce an
// empty result set for a query that asked for effectively no limit, which is
// the opposite of what was written.
func TestAtoiClampsAnOverflowingLimitToKeepAll(t *testing.T) {
	const maxInt = int(^uint(0) >> 1)

	if got := atoi("99999999999999999999999999"); got != maxInt {
		t.Fatalf("atoi(overflowing) = %d, want %d (keep all)", got, maxInt)
	}
	if got := atoi("0"); got != 0 {
		t.Fatalf("atoi(\"0\") = %d, want 0 — an explicit zero is not an overflow", got)
	}
	if got := atoi("42"); got != 42 {
		t.Fatalf("atoi(\"42\") = %d, want 42", got)
	}
}

// TestLimitOfAnOverflowingLiteralKeepsEveryObject is the same claim seen from
// the query, which is where it matters: the clamp must survive parsing into a
// stage that keeps the whole result set.
func TestLimitOfAnOverflowingLiteralKeepsEveryObject(t *testing.T) {
	pipeline := mustParse(t, "limit 99999999999999999999999999")

	limit, ok := pipeline.Stages[0].(ast.LimitStage)
	if !ok {
		t.Fatalf("got %T, want ast.LimitStage", pipeline.Stages[0])
	}
	if limit.N != int(^uint(0)>>1) {
		t.Fatalf("N = %d, want max int so that every object is kept", limit.N)
	}
}
