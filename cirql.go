// Package cirql is the public entry point to the cirql JSON pipeline language:
// Parse compiles a query string into a Pipeline, and Pipeline.Run executes it
// over a decoded JSON value. The language, evaluator, and stages live in the
// sub-packages; this package wires them together.
package cirql

import (
	dialect "github.com/gomatic/cirql/pkg/dialect/cirql"
	"github.com/gomatic/cirql/pipeline"
	"github.com/gomatic/cirql/stage"
	"github.com/gomatic/cirql/value"
)

// Pipeline is a compiled cirql query ready to run.
type Pipeline struct {
	stages []pipeline.Stage
}

// config holds Parse options.
type config struct {
	now stage.Clock
}

// Option configures Parse.
type Option func(*config)

// WithClock injects the clock the now() builtin reads (epoch seconds), making
// time-dependent queries deterministic in tests.
func WithClock(now func() int64) Option {
	return func(c *config) { c.now = now }
}

// Parse compiles a cirql query into a Pipeline. It returns the dialect's
// ErrParse on a syntax error, or stage.ErrStageUnsupported when the query uses a
// source stage this build does not execute.
func Parse(query string, opts ...Option) (Pipeline, error) {
	var cfg config
	for _, o := range opts {
		o(&cfg)
	}
	tree, err := dialect.Parse(query)
	if err != nil {
		return Pipeline{}, err
	}
	stages := make([]pipeline.Stage, 0, len(tree.Stages))
	for _, s := range tree.Stages {
		st, berr := stage.Build(s, cfg.now)
		if berr != nil {
			return Pipeline{}, berr
		}
		stages = append(stages, st)
	}
	return Pipeline{stages: stages}, nil
}

// Run executes the pipeline over a decoded JSON value, normalizing it into the
// initial result set and threading it through every stage.
func (p Pipeline) Run(in value.Value) ([]value.Value, error) {
	return pipeline.RunStages(p.stages, pipeline.Normalize(in))
}
