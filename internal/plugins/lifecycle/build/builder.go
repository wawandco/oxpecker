package build

import (
	"context"

	"github.com/wawandco/oxpecker/plugins"
)

// Builder interface allows to set the build steps to be run.
type Builder interface {
	plugins.Plugin
	Build(context.Context, string, []string) error
}
