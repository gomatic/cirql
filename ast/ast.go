// Package ast defines the cirql abstract syntax tree: the typed pipeline,
// stages, and expressions a parsed query produces. Nodes are pure data; the
// closed Stage and Expr sets are enforced by unexported tag methods.
package ast

import value "github.com/gomatic/go-json"

// Pipeline is a parsed cirql query: an ordered list of stages.
type Pipeline struct {
	Stages []Stage
}

// Stage is one step in a pipeline. The set is closed via isStage.
type Stage interface{ isStage() }

// Expr is an expression evaluated against the current object. Closed via isExpr.
type Expr interface{ isExpr() }

// Mapping is a key/expression pair in a map or flatMap stage.
type Mapping struct {
	Expr Expr
	Key  string
}

// ReduceOp names an aggregation in a reduce stage.
type ReduceOp string

const (
	// OpCount counts the result set.
	OpCount ReduceOp = "count"
	// OpSum sums the argument expression over the result set.
	OpSum ReduceOp = "sum"
	// OpMin takes the minimum of the argument expression.
	OpMin ReduceOp = "min"
	// OpMax takes the maximum of the argument expression.
	OpMax ReduceOp = "max"
	// OpAvg averages the argument expression.
	OpAvg ReduceOp = "avg"
	// OpFirst returns the first object.
	OpFirst ReduceOp = "first"
	// OpLast returns the last object.
	OpLast ReduceOp = "last"
	// OpGroupBy groups objects by the argument expression.
	OpGroupBy ReduceOp = "group_by"
	// OpCollect collects the argument expression from each object into a list.
	OpCollect ReduceOp = "collect"
)

// BinOp names a binary operator.
type BinOp string

const (
	// OpOr is logical ||.
	OpOr BinOp = "||"
	// OpAnd is logical &&.
	OpAnd BinOp = "&&"
	// OpEq is ==.
	OpEq BinOp = "=="
	// OpNe is !=.
	OpNe BinOp = "!="
	// OpGt is >.
	OpGt BinOp = ">"
	// OpLt is <.
	OpLt BinOp = "<"
	// OpGe is >=.
	OpGe BinOp = ">="
	// OpLe is <=.
	OpLe BinOp = "<="
	// OpAdd is +.
	OpAdd BinOp = "+"
	// OpSub is -.
	OpSub BinOp = "-"
	// OpMul is *.
	OpMul BinOp = "*"
	// OpDiv is /.
	OpDiv BinOp = "/"
	// OpMod is %.
	OpMod BinOp = "%"
)

// UnOp names a unary operator.
type UnOp string

const (
	// OpNot is logical negation.
	OpNot UnOp = "!"
	// OpNeg is arithmetic negation (unary minus).
	OpNeg UnOp = "-"
)

// Transform stages.

// MapStage transforms each object 1:1 into a new object.
type MapStage struct{ Mappings []Mapping }

// FlatMapStage transforms and flattens: list-valued mappings expand into
// multiple output objects.
type FlatMapStage struct{ Mappings []Mapping }

// FilterStage keeps objects for which Cond is truthy.
type FilterStage struct{ Cond Expr }

// ReduceStage aggregates the result set into a single object.
type ReduceStage struct {
	Arg Expr
	Op  ReduceOp
}

// SortStage reorders the result set by Key.
type SortStage struct {
	Key    Expr
	IsDesc bool
}

// LimitStage truncates the result set to at most N objects.
type LimitStage struct{ N int }

// UniqStage deduplicates the result set, by Key when non-nil else whole-object.
type UniqStage struct{ Key Expr }

// Source stages (parsed in core; executed in the sources sub-project).

// StdinStage reads the pipeline input as the result set.
type StdinStage struct{}

// FileStage reads JSON from a local file.
type FileStage struct{ Path string }

// HTTPStage fetches JSON over HTTP.
type HTTPStage struct{ URL Expr }

// QueryStage executes a GraphQL query.
type QueryStage struct{ Body string }

func (MapStage) isStage()     {}
func (FlatMapStage) isStage() {}
func (FilterStage) isStage()  {}
func (ReduceStage) isStage()  {}
func (SortStage) isStage()    {}
func (LimitStage) isStage()   {}
func (UniqStage) isStage()    {}
func (StdinStage) isStage()   {}
func (FileStage) isStage()    {}
func (HTTPStage) isStage()    {}
func (QueryStage) isStage()   {}

// Expressions.

// PathSegment is one step of a field access: a named field, or [] iteration.
type PathSegment struct {
	Name   string
	IsIter bool
}

// FieldAccess is .a.b[] etc.; an empty Path is the identity ".".
type FieldAccess struct{ Path []PathSegment }

// VarRef is a $name reference.
type VarRef struct{ Name string }

// BinaryExpr is L Op R.
type BinaryExpr struct {
	L  Expr
	R  Expr
	Op BinOp
}

// UnaryExpr is Op X.
type UnaryExpr struct {
	X  Expr
	Op UnOp
}

// FuncCall is name(args...).
type FuncCall struct {
	Name string
	Args []Expr
}

// Literal is a constant value.
type Literal struct{ V value.Value }

func (FieldAccess) isExpr() {}
func (VarRef) isExpr()      {}
func (BinaryExpr) isExpr()  {}
func (UnaryExpr) isExpr()   {}
func (FuncCall) isExpr()    {}
func (Literal) isExpr()     {}
