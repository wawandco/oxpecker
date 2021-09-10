package fix

import (
	"context"

	"github.com/wawandco/ox/plugins"
)

// Fixer interface is created for those commands that fill fix certain
// part of our code to match versions or compile correctly. Fixers are
// a good way to transition from old versions of the tools into newer
// ones
type Fixer interface {
	plugins.Plugin
	Fix(context.Context, string, []string) error
}
