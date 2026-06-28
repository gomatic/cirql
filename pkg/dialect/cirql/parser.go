// Package cirql is the covered seam over the generated ANTLR parser: it turns a
// cirql query string into a typed ast.Pipeline and any syntax error into the
// ErrParse sentinel. The generated lexer/parser live in src/grammar/cirql and
// are never edited by hand.
package cirql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gomatic/cirql/ast"
	g "github.com/gomatic/cirql/src/grammar/cirql"
	value "github.com/gomatic/go-json"
)

// Parse turns a cirql query into a Pipeline AST, or ErrParse on a syntax error.
func Parse(query string) (ast.Pipeline, error) {
	el := &errListener{}
	lexer := g.NewcirqlLexer(antlr.NewInputStream(query))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := g.NewcirqlParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(el)
	tree := parser.Pipeline()
	if el.err != nil {
		return ast.Pipeline{}, el.err
	}
	return (&builder{}).pipeline(tree.(*g.PipelineContext)), nil
}

// errListener converts an ANTLR syntax error into ErrParse with position.
type errListener struct {
	*antlr.DefaultErrorListener
	err error
}

// SyntaxError records the syntax error (last one wins) as an ErrParse wrap.
func (l *errListener) SyntaxError(_ antlr.Recognizer, _ any, line, col int, msg string, _ antlr.RecognitionException) {
	l.err = fmt.Errorf("%w at %d:%d: %s", ErrParse, line, col, msg)
}

// builder walks the parse tree into the AST. Every method is total: the parse
// tree is already valid, and number/string conversions are made lenient so no
// build step can fail.
type builder struct{}

func (b *builder) pipeline(c *g.PipelineContext) ast.Pipeline {
	stages := make([]ast.Stage, 0, len(c.AllStage()))
	for _, s := range c.AllStage() {
		stages = append(stages, b.stage(s.(*g.StageContext)))
	}
	return ast.Pipeline{Stages: stages}
}

func (b *builder) stage(c *g.StageContext) ast.Stage {
	if src := c.SourceStage(); src != nil {
		return b.sourceStage(src.(*g.SourceStageContext))
	}
	return b.transformStage(c.TransformStage().(*g.TransformStageContext))
}

func (b *builder) sourceStage(c *g.SourceStageContext) ast.Stage {
	if c.StdinStage() != nil {
		return ast.StdinStage{}
	}
	if s := c.FileStage(); s != nil {
		return ast.FileStage{Path: unquote(s.(*g.FileStageContext).STRING().GetText())}
	}
	if s := c.HttpStage(); s != nil {
		return b.httpStage(s.(*g.HttpStageContext))
	}
	return ast.QueryStage{Body: c.QueryStage().(*g.QueryStageContext).QueryBody().GetText()}
}

func (b *builder) httpStage(c *g.HttpStageContext) ast.Stage {
	if v := c.Variable(); v != nil {
		return ast.HTTPStage{URL: b.variable(v.(*g.VariableContext))}
	}
	return ast.HTTPStage{URL: ast.Literal{V: unquote(c.STRING().GetText())}}
}

func (b *builder) transformStage(c *g.TransformStageContext) ast.Stage {
	if s := c.MapStage(); s != nil {
		return ast.MapStage{Mappings: b.mappings(s.(*g.MapStageContext).AllMapping())}
	}
	if s := c.FlatMapStage(); s != nil {
		return ast.FlatMapStage{Mappings: b.mappings(s.(*g.FlatMapStageContext).AllMapping())}
	}
	if s := c.FilterStage(); s != nil {
		return ast.FilterStage{Cond: b.expr(s.(*g.FilterStageContext).Expr())}
	}
	if s := c.ReduceStage(); s != nil {
		return b.reduceStage(s.(*g.ReduceStageContext))
	}
	if s := c.SortStage(); s != nil {
		return b.sortStage(s.(*g.SortStageContext))
	}
	if s := c.LimitStage(); s != nil {
		return ast.LimitStage{N: atoi(s.(*g.LimitStageContext).INT().GetText())}
	}
	return b.uniqStage(c.UniqStage().(*g.UniqStageContext))
}

func (b *builder) mappings(all []g.IMappingContext) []ast.Mapping {
	out := make([]ast.Mapping, 0, len(all))
	for _, m := range all {
		mc := m.(*g.MappingContext)
		out = append(out, ast.Mapping{Key: mc.NAME().GetText(), Expr: b.expr(mc.Expr())})
	}
	return out
}

func (b *builder) reduceStage(c *g.ReduceStageContext) ast.Stage {
	stage := ast.ReduceStage{Op: ast.ReduceOp(c.ReduceOp().GetText())}
	if e := c.Expr(); e != nil {
		stage.Arg = b.expr(e)
	}
	return stage
}

func (b *builder) sortStage(c *g.SortStageContext) ast.Stage {
	return ast.SortStage{Key: b.expr(c.Expr()), Desc: c.DESC() != nil}
}

func (b *builder) uniqStage(c *g.UniqStageContext) ast.Stage {
	stage := ast.UniqStage{}
	if e := c.Expr(); e != nil {
		stage.Key = b.expr(e)
	}
	return stage
}

func (b *builder) expr(ctx g.IExprContext) ast.Expr {
	switch c := ctx.(type) {
	case *g.UnaryExprContext:
		return b.unary(c)
	case *g.MulExprContext:
		return b.binary(c.AllExpr(), childOp(c))
	case *g.AddExprContext:
		return b.binary(c.AllExpr(), childOp(c))
	case *g.CmpExprContext:
		return b.binary(c.AllExpr(), childOp(c))
	case *g.AndExprContext:
		return b.binary(c.AllExpr(), childOp(c))
	case *g.OrExprContext:
		return b.binary(c.AllExpr(), childOp(c))
	case *g.ParenExprContext:
		return b.expr(c.Expr())
	case *g.CallExprContext:
		return b.funcCall(c.FuncCall().(*g.FuncCallContext))
	case *g.FieldExprContext:
		return b.fieldAccess(c.FieldAccess().(*g.FieldAccessContext))
	case *g.VarExprContext:
		return b.variable(c.Variable().(*g.VariableContext))
	default:
		return b.literal(ctx.(*g.LitExprContext))
	}
}

func (b *builder) unary(c *g.UnaryExprContext) ast.Expr {
	x := b.expr(c.Expr())
	if c.NOT() != nil {
		return ast.UnaryExpr{Op: ast.OpNot, X: x}
	}
	return ast.UnaryExpr{Op: ast.OpNeg, X: x}
}

func (b *builder) binary(all []g.IExprContext, op ast.BinOp) ast.Expr {
	return ast.BinaryExpr{Op: op, L: b.expr(all[0]), R: b.expr(all[1])}
}

func (b *builder) funcCall(c *g.FuncCallContext) ast.Expr {
	args := make([]ast.Expr, 0, len(c.AllExpr()))
	for _, e := range c.AllExpr() {
		args = append(args, b.expr(e))
	}
	return ast.FuncCall{Name: c.NAME().GetText(), Args: args}
}

func (b *builder) fieldAccess(c *g.FieldAccessContext) ast.Expr {
	fa := ast.FieldAccess{}
	for _, ps := range c.AllPathSeg() {
		psc := ps.(*g.PathSegContext)
		if psc.NAME() != nil {
			fa.Path = append(fa.Path, ast.PathSegment{Name: psc.NAME().GetText()})
			continue
		}
		fa.Path = append(fa.Path, ast.PathSegment{Iter: true})
	}
	return fa
}

func (b *builder) variable(c *g.VariableContext) ast.Expr {
	return ast.VarRef{Name: c.NAME().GetText()}
}

func (b *builder) literal(c *g.LitExprContext) ast.Expr {
	return ast.Literal{V: literalValue(c.Literal().(*g.LiteralContext))}
}

// childOp reads the operator token of a binary expression alternative (child 1
// is the operator between the two operand expressions).
func childOp(ctx antlr.ParserRuleContext) ast.BinOp {
	return ast.BinOp(ctx.GetChild(1).(antlr.TerminalNode).GetText())
}

// literalValue converts a literal token into a value.Value.
func literalValue(c *g.LiteralContext) value.Value {
	if t := c.STRING(); t != nil {
		return unquote(t.GetText())
	}
	if t := c.FLOAT(); t != nil {
		return parseFloat(t.GetText())
	}
	if t := c.INT(); t != nil {
		return parseInt(t.GetText())
	}
	if c.TRUE() != nil {
		return true
	}
	if c.FALSE() != nil {
		return false
	}
	return nil
}

// parseInt parses a decimal integer, falling back to float on overflow.
func parseInt(s string) value.Value {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		f, _ := strconv.ParseFloat(s, 64)
		return f
	}
	return n
}

// parseFloat parses a float, taking the rounded value even on range overflow.
func parseFloat(s string) value.Value {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// atoi parses a non-negative integer token (guaranteed digits by the lexer);
// an overflowing limit clamps to the max int, i.e. "keep all".
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return int(^uint(0) >> 1)
	}
	return n
}

// unquote strips the surrounding quotes of a STRING token and resolves its
// escapes leniently (an unknown escape yields the escaped character verbatim).
func unquote(s string) string {
	inner := s[1 : len(s)-1]
	var b strings.Builder
	for i := 0; i < len(inner); i++ {
		if inner[i] == '\\' && i+1 < len(inner) {
			i++
			b.WriteByte(unescape(inner[i]))
			continue
		}
		b.WriteByte(inner[i])
	}
	return b.String()
}

// unescape maps a backslash-escaped byte to its value; unknown escapes pass the
// byte through unchanged.
func unescape(c byte) byte {
	switch c {
	case 'n':
		return '\n'
	case 't':
		return '\t'
	case 'r':
		return '\r'
	}
	return c
}
