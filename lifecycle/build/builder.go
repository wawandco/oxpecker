package build

import (
	"context"

	"github.com/wawandco/ox/plugins"
)

// Builder interface allows to set the build steps to be run.
type Builder interface {
	plugins.Plugin
	Build(context.Context, string, []string) error
}
