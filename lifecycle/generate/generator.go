package generate

import (
	"context"

	"github.com/wawandco/oxpecker/plugins"
)

// Generator allows to identify those plugins that are
// generators.
type Generator interface {
	plugins.Plugin
	Generate(context.Context, string, []string) error
}
