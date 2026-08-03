package eval

import errs "github.com/gomatic/go-error"

// Error is the sentinel-error type for the eval package — an alias of
// [errs.Const] from github.com/gomatic/go-error, which owns the mechanism
// (Error, With). The alias keeps every existing `eval.Error` reference and
// errors.Is match compiling and behaving identically.
type Error = errs.Const

const (
	// ErrUnknownFunc is returned when a called builtin does not exist.
	ErrUnknownFunc Error = "eval: unknown function"
	// ErrArity is returned when a builtin gets the wrong number of arguments.
	ErrArity Error = "eval: wrong argument count"
	// ErrType is returned when an operation gets an unsupported value type.
	ErrType Error = "eval: type error"
	// ErrNilExpr is returned when Eval is handed a nil expression — a pipeline
	// assembled programmatically with a missing Cond/Key/Expr. A parsed query
	// never reaches it: every expression the parser builds is non-nil, and the
	// two slots the grammar lets a query leave empty (a bare uniq's Key, a
	// count/first/last reduce's Arg) are guarded by their stage executor
	// before Eval is called.
	ErrNilExpr Error = "eval: nil expression"
)
