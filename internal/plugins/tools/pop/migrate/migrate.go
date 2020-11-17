package migrate

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gobuffalo/pop/v5"
	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Command    = (*Plugin)(nil)
	_ plugins.Subcommand = (*Plugin)(nil)
	_ plugins.FlagParser = (*Plugin)(nil)

	ErrCouldNotFindConnection = errors.New("could not find connection by name")
)

const (
	migrateUp   = "UP"
	migrateDown = "DOWN"
)

type Plugin struct {
	migrationPath  string
	connectionName string
	configFile     string

	// direction could be UP or DOWN, defaults to UP
	direction string

	// steps is the number of migrations to apply
	steps int
}

func (m *Plugin) Name() string {
	return "pop/migrate"
}

func (m *Plugin) SubcommandName() string {
	return "migrate"
}

func (m *Plugin) ParseFlags(args []string) error {
	m.migrationPath = "migrations"
	m.connectionName = "development"
	m.direction = "UP"
	m.configFile = "config/database.yml"
	m.steps = 1

	return nil
}

// Run will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (m *Plugin) Run(ctx context.Context, root string, args []string) error {
	cb, err := os.OpenFile(m.configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	err = pop.LoadFrom(cb)
	if err != nil {
		return err
	}

	conn := pop.Connections[m.connectionName]
	fmt.Println(m.connectionName)
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewFileMigrator(m.migrationPath, conn)
	if err != nil {
		return err
	}

	if m.direction == migrateUp {
		return mig.Up()
	}

	return mig.Down(m.steps)
}
