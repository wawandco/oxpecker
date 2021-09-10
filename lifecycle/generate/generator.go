package generate

import (
	"context"

	"github.com/wawandco/ox/plugins"
)

// Generator allows to identify those plugins that are
// generators.
type Generator interface {
	plugins.Plugin
	InvocationName() string
	Generate(context.Context, string, []string) error
}

// After generator is something that runs after generators
// are executed.
type AfterGenerator interface {
	// AfterGenerate receives the context and other params so it can determine if should
	// run or not.
	plugins.Plugin
	AfterGenerate(context.Context, string, []string) error
}
