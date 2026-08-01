// Package cirql is the covered seam over the generated ANTLR parser: it turns a
// cirql query string into a typed ast.Pipeline and any syntax error into the
// ErrParse sentinel. The generated lexer/parser live in src/grammar/cirql and
// are never edited by hand.
package cirql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/gomatic/cirql/ast"
	g "github.com/gomatic/cirql/src/grammar/cirql"
)

// Query is the raw text of a cirql query.
type Query string

// MaxQueryBytes bounds a query's length. The generated parser recurses once per
// nesting level, so an unbounded query could overflow the goroutine stack (an
// uncatchable fatal error). 256 KiB caps nesting far below that limit while
// dwarfing any real query — query text is DSL, not data (data arrives on
// stdin), so this never constrains legitimate use.
const MaxQueryBytes = 256 * 1024

// Parse turns a cirql query into a Pipeline AST, or ErrParse on a syntax error.
// A query longer than MaxQueryBytes is rejected with ErrQueryTooLarge before
// parsing, so adversarial deep nesting cannot overflow the stack.
func Parse(query Query) (ast.Pipeline, error) {
	if len(query) > MaxQueryBytes {
		return ast.Pipeline{}, ErrQueryTooLarge.With(nil, fmt.Sprintf("%d bytes exceeds %d", len(query), MaxQueryBytes))
	}
	var parseErr error
	el := errListener{err: &parseErr}
	lexer := g.NewcirqlLexer(antlr.NewInputStream(string(query)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := g.NewcirqlParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(el)
	tree := parser.Pipeline()
	if parseErr != nil {
		return ast.Pipeline{}, parseErr
	}
	return builder{}.pipeline(tree), nil
}

// errListener converts an ANTLR syntax error into ErrParse with position. It is
// a value handle over a shared error slot (the pointer field), so it satisfies
// antlr.ErrorListener with a value receiver.
type errListener struct {
	*antlr.DefaultErrorListener
	err *error
}

// SyntaxError records the syntax error (last one wins) as an ErrParse wrap.
func (l errListener) SyntaxError(_ antlr.Recognizer, _ any, line, col int, msg string, _ antlr.RecognitionException) {
	*l.err = ErrParse.With(nil, fmt.Sprintf("at %d:%d: %s", line, col, msg))
}

// builder walks the parse tree into the AST. Every method is total: the parse
// tree is already valid, and number/string conversions are made lenient so no
// build step can fail. Methods take the generated I*Context interfaces so the
// tree is traversed by contract, not by concrete pointer.
type builder struct{}

func (b builder) pipeline(c g.IPipelineContext) ast.Pipeline {
	stages := make([]ast.Stage, 0, len(c.AllStage()))
	for _, s := range c.AllStage() {
		stages = append(stages, b.stage(s))
	}
	return ast.Pipeline{Stages: stages}
}

func (b builder) stage(c g.IStageContext) ast.Stage {
	if src := c.SourceStage(); src != nil {
		return b.sourceStage(src)
	}
	return b.transformStage(c.TransformStage())
}

func (b builder) sourceStage(c g.ISourceStageContext) ast.Stage {
	if c.StdinStage() != nil {
		return ast.StdinStage{}
	}
	if s := c.FileStage(); s != nil {
		return ast.FileStage{Path: unquote(tokenText(s.STRING().GetText()))}
	}
	if s := c.HttpStage(); s != nil {
		return b.httpStage(s)
	}
	return ast.QueryStage{Body: c.QueryStage().QueryBody().GetText()}
}

func (b builder) httpStage(c g.IHttpStageContext) ast.Stage {
	if v := c.Variable(); v != nil {
		return ast.HTTPStage{URL: b.variable(v)}
	}
	return ast.HTTPStage{URL: ast.Literal{V: unquote(tokenText(c.STRING().GetText()))}}
}

func (b builder) transformStage(c g.ITransformStageContext) ast.Stage {
	if s := c.MapStage(); s != nil {
		return ast.MapStage{Mappings: b.mappings(s.AllMapping())}
	}
	if s := c.FlatMapStage(); s != nil {
		return ast.FlatMapStage{Mappings: b.mappings(s.AllMapping())}
	}
	if s := c.FilterStage(); s != nil {
		return ast.FilterStage{Cond: b.expr(s.Expr())}
	}
	if s := c.ReduceStage(); s != nil {
		return b.reduceStage(s)
	}
	if s := c.SortStage(); s != nil {
		return b.sortStage(s)
	}
	if s := c.LimitStage(); s != nil {
		return ast.LimitStage{N: atoi(tokenText(s.INT().GetText()))}
	}
	return b.uniqStage(c.UniqStage())
}

func (b builder) mappings(all []g.IMappingContext) []ast.Mapping {
	out := make([]ast.Mapping, 0, len(all))
	for _, m := range all {
		out = append(out, ast.Mapping{Key: m.FieldName().GetText(), Expr: b.expr(m.Expr())})
	}
	return out
}

func (b builder) reduceStage(c g.IReduceStageContext) ast.Stage {
	stage := ast.ReduceStage{Op: ast.ReduceOp(c.ReduceOp().GetText())}
	if e := c.Expr(); e != nil {
		stage.Arg = b.expr(e)
	}
	return stage
}

func (b builder) sortStage(c g.ISortStageContext) ast.Stage {
	return ast.SortStage{Key: b.expr(c.Expr()), IsDesc: c.DESC() != nil}
}

func (b builder) uniqStage(c g.IUniqStageContext) ast.Stage {
	stage := ast.UniqStage{}
	if e := c.Expr(); e != nil {
		stage.Key = b.expr(e)
	}
	return stage
}
