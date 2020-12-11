package migrate

import "context"

// Migrator type allows to identify migration runner plugins.
type Migrator interface {
	CommandName() string
	RunMigrations(ctx context.Context, root string, args []string) error
}
