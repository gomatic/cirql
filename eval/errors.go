package eval

// Error is the sentinel error type for the eval package.
type Error string

// Error renders the sentinel message.
func (e Error) Error() string { return string(e) }

const (
	// ErrUnknownFunc is returned when a called builtin does not exist.
	ErrUnknownFunc Error = "eval: unknown function"
	// ErrArity is returned when a builtin gets the wrong number of arguments.
	ErrArity Error = "eval: wrong argument count"
	// ErrType is returned when an operation gets an unsupported value type.
	ErrType Error = "eval: type error"
)
