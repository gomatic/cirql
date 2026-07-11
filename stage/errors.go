package stage

import errs "github.com/gomatic/go-error"

// Error is the sentinel-error type for the stage package — an alias of
// [errs.Const] from github.com/gomatic/go-error, which owns the mechanism
// (Error, With). The alias keeps every existing `stage.Error` reference and
// errors.Is match compiling and behaving identically.
type Error = errs.Const

// ErrStageUnsupported is returned when building a source stage (file/http/query)
// that cirql-core does not execute; the sources sub-project provides these.
const ErrStageUnsupported Error = "cirql: source stage not supported in core"
