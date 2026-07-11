package cirql

import errs "github.com/gomatic/go-error"

// ErrParse is returned for any lexer or parser syntax error; the wrap adds the
// source position and the underlying message.
const ErrParse errs.Const = "cirql: parse error"

// ErrQueryTooLarge is returned when a query exceeds MaxQueryBytes. The bound
// exists because the generated recursive-descent parser recurses per nesting
// level, so a deeply nested adversarial query (e.g. hundreds of thousands of
// parentheses) would overflow the goroutine stack — an uncatchable fatal error,
// not a panic. Bounding the query length caps nesting depth well below that,
// and a cirql query is bounded DSL text (the DATA arrives on stdin, never in
// the query), so the limit never constrains a real query.
const ErrQueryTooLarge errs.Const = "cirql: query too large"
