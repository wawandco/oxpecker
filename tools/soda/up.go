package soda

import (
	"context"

	"github.com/gobuffalo/pop/v5"
)

// Run will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (mu *Command) RunUp() error {
	pop.SetLogger(mu.Log)

	conn := pop.Connections[mu.connectionName]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewMigrationBox(mu.migrations, conn)
	if err != nil {
		return err
	}

	_, err = mig.UpTo(mu.steps)
	return err
}

func (mu *Command) RunBeforeTest(ctx context.Context, root string, args []string) error {
	pop.SetLogger(mu.Log)

	conn := pop.Connections["test"]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewMigrationBox(mu.migrations, conn)
	if err != nil {
		return err
	}

	err = mig.Up()
	return err
}
