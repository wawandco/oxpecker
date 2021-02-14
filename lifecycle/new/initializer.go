package new

import (
	"context"
	"sync"
)

// Initializer is intended to initialize applications,
// things like generating files or running commands.
type Initializer interface {
	// Initialize receives the context and the root folder where
	// the application is being initialized.
	Initialize(context.Context, *sync.Map) error
}
