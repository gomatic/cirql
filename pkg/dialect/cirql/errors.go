package cirql

// Error is the sentinel error type for the cirql dialect wrapper.
type Error string

// Error renders the sentinel message.
func (e Error) Error() string { return string(e) }

// ErrParse is returned for any lexer or parser syntax error; the wrap adds the
// source position and the underlying message.
const ErrParse Error = "cirql: parse error"
