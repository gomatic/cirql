package cirql

import errs "github.com/gomatic/go-error"

// ErrParse is returned for any lexer or parser syntax error; the wrap adds the
// source position and the underlying message.
const ErrParse errs.Const = "cirql: parse error"
