package soda

import (
	"github.com/gobuffalo/pop/v5"
)

// RunDown will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (mu *Command) RunDown() error {
	pop.SetLogger(mu.Log)
	conn := pop.Connections[mu.connectionName]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewMigrationBox(mu.migrations, conn)
	if err != nil {
		return err
	}

	if mu.steps == 0 {
		mu.steps = 1
	}

	return mig.Down(mu.steps)
}
