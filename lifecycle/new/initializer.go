package new

import (
	"context"

	"github.com/wawandco/oxpecker/plugins"
)

// Initializer is intended to initialize applications,
// things like generating files or running commands.
type Initializer interface {
	// Initializers may require to use the passed flags.
	plugins.FlagParser

	// Initialize receives the context and the root folder where
	// the application is being initialized.
	Initialize(context.Context, string, []string) error
}
