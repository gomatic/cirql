package cirql

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	value "github.com/gomatic/go-json"

	"github.com/gomatic/cirql/ast"
	g "github.com/gomatic/cirql/src/grammar/cirql"
)

// Building expression nodes from the generated parse tree, and the token
// decoders they rest on. Split from the stage builders because the two halves
// change for different reasons: a new STAGE is a grammar rule and a new
// ast.Stage, while a new operator or literal form lands entirely here.

func (b builder) expr(ctx g.IExprContext) ast.Expr {
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
		return b.funcCall(c.FuncCall())
	case *g.FieldExprContext:
		return b.fieldAccess(c.FieldAccess())
	case *g.VarExprContext:
		return b.variable(c.Variable())
	default:
		return b.literal(ctx.(*g.LitExprContext))
	}
}

func (b builder) unary(c *g.UnaryExprContext) ast.Expr {
	x := b.expr(c.Expr())
	if c.NOT() != nil {
		return ast.UnaryExpr{Op: ast.OpNot, X: x}
	}
	return ast.UnaryExpr{Op: ast.OpNeg, X: x}
}

func (b builder) binary(all []g.IExprContext, op ast.BinOp) ast.Expr {
	return ast.BinaryExpr{Op: op, L: b.expr(all[0]), R: b.expr(all[1])}
}

func (b builder) funcCall(c g.IFuncCallContext) ast.Expr {
	args := make([]ast.Expr, 0, len(c.AllExpr()))
	for _, e := range c.AllExpr() {
		args = append(args, b.expr(e))
	}
	return ast.FuncCall{Name: c.NAME().GetText(), Args: args}
}

func (b builder) fieldAccess(c g.IFieldAccessContext) ast.Expr {
	fa := ast.FieldAccess{}
	if head := c.FieldName(); head != nil {
		fa.Path = append(fa.Path, ast.PathSegment{Name: head.GetText()})
	} else if c.LBRACK() != nil {
		fa.Path = append(fa.Path, ast.PathSegment{IsIter: true})
	}
	for _, ps := range c.AllPathSeg() {
		if fn := ps.FieldName(); fn != nil {
			fa.Path = append(fa.Path, ast.PathSegment{Name: fn.GetText()})
			continue
		}
		fa.Path = append(fa.Path, ast.PathSegment{IsIter: true})
	}
	return fa
}

func (b builder) variable(c g.IVariableContext) ast.Expr {
	return ast.VarRef{Name: c.NAME().GetText()}
}

func (b builder) literal(c *g.LitExprContext) ast.Expr {
	return ast.Literal{V: literalValue(c.Literal())}
}

// childOp reads the operator token of a binary expression alternative (child 1
// is the operator between the two operand expressions).
func childOp(ctx antlr.ParserRuleContext) ast.BinOp {
	return ast.BinOp(ctx.GetChild(1).(antlr.TerminalNode).GetText())
}

// literalValue converts a literal token into a value.Value.
func literalValue(c g.ILiteralContext) value.Value {
	if t := c.STRING(); t != nil {
		return unquote(tokenText(t.GetText()))
	}
	if t := c.FLOAT(); t != nil {
		return parseFloat(tokenText(t.GetText()))
	}
	if t := c.INT(); t != nil {
		return parseInt(tokenText(t.GetText()))
	}
	if c.TRUE() != nil {
		return true
	}
	if c.FALSE() != nil {
		return false
	}
	return nil
}

// tokenText is the raw source text of a lexer token (STRING, FLOAT, or INT).
type tokenText string

// parseInt parses a decimal integer, falling back to float on overflow.
func parseInt(s tokenText) value.Value {
	n, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		f, _ := strconv.ParseFloat(string(s), 64)
		return f
	}
	return n
}

// parseFloat parses a float, taking the rounded value even on range overflow.
func parseFloat(s tokenText) value.Value {
	f, _ := strconv.ParseFloat(string(s), 64)
	return f
}

// atoi parses a non-negative integer token (guaranteed digits by the lexer);
// an overflowing limit clamps to the max int, i.e. "keep all".
func atoi(s tokenText) int {
	n, err := strconv.Atoi(string(s))
	if err != nil {
		return int(^uint(0) >> 1)
	}
	return n
}

// unquote strips the surrounding quotes of a STRING token and resolves its
// escapes leniently (an unknown escape yields the escaped character verbatim).
func unquote(s tokenText) string {
	inner := string(s)[1 : len(string(s))-1]
	var b strings.Builder
	for i := 0; i < len(inner); i++ {
		if inner[i] == '\\' && i+1 < len(inner) {
			i++
			_ = b.WriteByte(unescape(escapedByte(inner[i])))
			continue
		}
		_ = b.WriteByte(inner[i])
	}
	return b.String()
}

// escapedByte is the byte following a backslash in a STRING token's escape sequence.
type escapedByte byte

// unescape maps a backslash-escaped byte to its value; unknown escapes pass the
// byte through unchanged.
func unescape(c escapedByte) byte {
	switch byte(c) {
	case 'n':
		return '\n'
	case 't':
		return '\t'
	case 'r':
		return '\r'
	}
	return byte(c)
}
