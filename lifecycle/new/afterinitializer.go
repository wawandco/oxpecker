package new

import (
	"context"
)

// AfterInitializer is intended to initialize applications,
// things that should happen at the end of the application
// creation process.
type AfterInitializer interface {
	// AfterInitialize receives the context and the root folder where
	// the application is being initialized.
	AfterInitialize(context.Context, Options) error
}
