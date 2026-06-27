package stage

// Error is the sentinel error type for the stage package.
type Error string

// Error renders the sentinel message.
func (e Error) Error() string { return string(e) }

// ErrStageUnsupported is returned when building a source stage (file/http/query)
// that cirql-core does not execute; the sources sub-project provides these.
const ErrStageUnsupported Error = "cirql: source stage not supported in core"
