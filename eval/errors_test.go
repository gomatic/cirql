package eval

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/gomatic/cirql/ast"
	dialect "github.com/gomatic/cirql/pkg/dialect/cirql"
)

// parserCorpus spans every stage form and every expression form the cirql
// grammar can produce. ErrNilExpr's claim is about the PARSER, so a form left
// out of this corpus is a form the claim is not tested for.
var parserCorpus = []string{
	`map { a: .a, b: .b.c[] }`,
	`flatMap { a: .a[] }`,
	`filter .stars > 1000 && !(.name == "go")`,
	`filter -.a + 1 * 2 / 3 - 4 % 5 >= 6.5`,
	`filter length(.a) != 0 || $v <= true`,
	`filter .a < false`,
	`sort .a asc`,
	`sort .a desc`,
	`limit 10`,
	`uniq`,
	`uniq .a`,
	`reduce count`,
	`reduce first`,
	`reduce last`,
	`reduce sum(.a)`,
	`reduce min(.a)`,
	`reduce max(.a)`,
	`reduce avg(.a)`,
	`reduce group_by(.a)`,
	`reduce collect(.a)`,
	`stdin`,
	`file "x.json"`,
	`http "http://h"`,
	`http $u`,
	`query { a b }`,
	`{ a b }`,
}

// omittableSlots are the only two expression slots the grammar lets a query
// leave empty: a bare `uniq` has no key, and `reduce count`/`first`/`last` has
// no argument. Both are guarded by their stage executor before Eval is called,
// which is why an empty slot is not a nil expression reaching Eval.
var omittableSlots = []string{"ReduceStage.Arg", "UniqStage.Key"}

// exprSlots is what one walk of a parsed pipeline found: every expression the
// parser BUILT, and the names of the optional slots it left empty.
type exprSlots struct {
	built   []ast.Expr
	omitted []string
}

// merge concatenates two walks.
func (s exprSlots) merge(other exprSlots) exprSlots {
	return exprSlots{
		built:   append(s.built, other.built...),
		omitted: append(s.omitted, other.omitted...),
	}
}

// TestErrNilExprIsUnreachableFromAParsedQuery names ErrNilExpr's documented
// claim: a parsed query never hands Eval a nil expression. That matters because
// a nil expression is the one input Eval cannot evaluate — if the parser could
// emit one, an ordinary user query would fail with an internal-sounding
// "nil expression" error instead of running.
//
// The walk asserts the claim from both sides: every expression the parser built
// is non-nil AND evaluates without reporting ErrNilExpr, and the only slots left
// empty are the two the grammar permits — each of which the corpus must actually
// reach, so the second assertion cannot pass vacuously.
func TestErrNilExprIsUnreachableFromAParsedQuery(t *testing.T) {
	t.Parallel()

	found := corpusSlots(t)

	if len(found.built) == 0 {
		t.Fatal("the corpus produced no expressions; the walk proves nothing")
	}
	for i, expr := range found.built {
		if _, err := Eval(expr, Env{}); errors.Is(err, ErrNilExpr) {
			t.Fatalf("parsed expression %d (%T) reported ErrNilExpr", i, expr)
		}
	}
	for _, slot := range found.omitted {
		if !slices.Contains(omittableSlots, slot) {
			t.Fatalf("the parser left %s empty, which no stage executor guards", slot)
		}
	}
	for _, slot := range omittableSlots {
		if !slices.Contains(found.omitted, slot) {
			t.Fatalf("no query in the corpus leaves %s empty; the guard it needs is untested", slot)
		}
	}
}

// TestErrNilExprIsReportedInsteadOfAPanic names ErrNilExpr's other half: the
// nil expression a programmatically assembled pipeline CAN hold is reported,
// not dereferenced.
func TestErrNilExprIsReportedInsteadOfAPanic(t *testing.T) {
	t.Parallel()

	_, err := Eval(nil, Env{})
	if !errors.Is(err, ErrNilExpr) {
		t.Errorf("got %v, want ErrNilExpr", err)
	}
}

// corpusSlots parses every query in parserCorpus and walks the resulting
// pipelines, failing on the first required expression the parser left nil.
func corpusSlots(t *testing.T) exprSlots {
	t.Helper()
	found := exprSlots{}
	for _, query := range parserCorpus {
		parsed, err := dialect.Parse(dialect.Query(query))
		if err != nil {
			t.Fatalf("Parse(%q): %v", query, err)
		}
		found = found.merge(pipelineSlots(t, parsed))
	}
	return found
}

// pipelineSlots walks every stage of one parsed pipeline.
func pipelineSlots(t *testing.T, parsed ast.Pipeline) exprSlots {
	t.Helper()
	found := exprSlots{}
	for _, stage := range parsed.Stages {
		found = found.merge(stageSlots(t, stage))
	}
	return found
}

// stageSlots walks the expression slots of one stage. An unhandled stage type
// fails the test rather than being skipped: a new ast.Stage must be added here,
// or the claim silently stops covering it.
func stageSlots(t *testing.T, stage ast.Stage) exprSlots {
	t.Helper()
	switch s := stage.(type) {
	case ast.MapStage:
		return mappingSlots(t, "MapStage.Mappings", s.Mappings)
	case ast.FlatMapStage:
		return mappingSlots(t, "FlatMapStage.Mappings", s.Mappings)
	case ast.FilterStage:
		return requiredSlot(t, "FilterStage.Cond", s.Cond)
	case ast.SortStage:
		return requiredSlot(t, "SortStage.Key", s.Key)
	case ast.HTTPStage:
		return requiredSlot(t, "HTTPStage.URL", s.URL)
	case ast.ReduceStage:
		return optionalSlot(t, "ReduceStage.Arg", s.Arg)
	case ast.UniqStage:
		return optionalSlot(t, "UniqStage.Key", s.Key)
	case ast.LimitStage, ast.StdinStage, ast.FileStage, ast.QueryStage:
		return exprSlots{}
	}
	t.Fatalf("unhandled ast.Stage %T; extend this walk or the claim stops covering it", stage)
	return exprSlots{}
}

// mappingSlots walks the expression of every mapping in a map/flatMap stage.
func mappingSlots(t *testing.T, where string, mappings []ast.Mapping) exprSlots {
	t.Helper()
	found := exprSlots{}
	for _, mapping := range mappings {
		found = found.merge(requiredSlot(t, where+"."+mapping.Key, mapping.Expr))
	}
	return found
}

// requiredSlot records an expression the grammar always fills, failing when the
// parser left it nil.
func requiredSlot(t *testing.T, where string, expr ast.Expr) exprSlots {
	t.Helper()
	if expr == nil {
		t.Fatalf("the parser left %s nil, which reaches Eval as a nil expression", where)
	}
	return exprSlots{built: []ast.Expr{expr}}.merge(childSlots(t, where, expr))
}

// optionalSlot records an expression the grammar may omit, noting the omission
// instead of failing.
func optionalSlot(t *testing.T, where string, expr ast.Expr) exprSlots {
	t.Helper()
	if expr == nil {
		return exprSlots{omitted: []string{where}}
	}
	return requiredSlot(t, where, expr)
}

// childSlots walks an expression's operands. An unhandled expression type fails
// the test for the same reason an unhandled stage does.
func childSlots(t *testing.T, where string, expr ast.Expr) exprSlots {
	t.Helper()
	switch e := expr.(type) {
	case ast.BinaryExpr:
		return requiredSlot(t, where+".L", e.L).merge(requiredSlot(t, where+".R", e.R))
	case ast.UnaryExpr:
		return requiredSlot(t, where+".X", e.X)
	case ast.FuncCall:
		return argSlots(t, where, e.Args)
	case ast.FieldAccess, ast.VarRef, ast.Literal:
		return exprSlots{}
	}
	t.Fatalf("unhandled ast.Expr %T at %s; extend this walk or the claim stops covering it", expr, where)
	return exprSlots{}
}

// argSlots walks the arguments of a function call.
func argSlots(t *testing.T, where string, args []ast.Expr) exprSlots {
	t.Helper()
	found := exprSlots{}
	for i, arg := range args {
		found = found.merge(requiredSlot(t, fmt.Sprintf("%s.Args[%d]", where, i), arg))
	}
	return found
}
