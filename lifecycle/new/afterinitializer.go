package new

import (
	"context"

	"github.com/wawandco/oxpecker/plugins"
)

// AfterInitializer is intended to initialize applications,
// things that should happen at the end of the application
// creation process.
type AfterInitializer interface {
	// AfterInitializers may require to use the passed flags.
	plugins.FlagParser

	// AfterInitialize receives the context and the root folder where
	// the application is being initialized.
	AfterInitialize(context.Context, string, []string) error
}
