package migrate

import (
	"context"
	"errors"
	"os"

	"github.com/gobuffalo/pop/v5"
	"github.com/paganotoni/oxpecker/internal/plugins"
	"github.com/spf13/pflag"
)

var (
	_ plugins.Command    = (*Plugin)(nil)
	_ plugins.Subcommand = (*Plugin)(nil)
	_ plugins.FlagParser = (*Plugin)(nil)

	migrateUp                 = "UP"
	ErrCouldNotFindConnection = errors.New("could not find connection by name")
)

type Plugin struct {
	migrationPath  string
	connectionName string
	configFile     string

	// direction could be UP or DOWN, defaults to UP
	direction string

	// steps is the number of migrations to apply
	steps int

	flags *pflag.FlagSet
}

//HelpText resturns the help Text of build function
func (m Plugin) HelpText() string {
	return `Runs migrations on the current folder, it looks for the
	migrations folder and attempt to run the migrations using internal
	pop tooling`
}

func (m *Plugin) Name() string {
	return "pop/migrate"
}

func (m *Plugin) SubcommandName() string {
	return "migrate"
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
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewFileMigrator(m.migrationPath, conn)
	if err != nil {
		return err
	}

	if m.direction == migrateUp || m.direction == "" {
		return mig.Up()
	}

	return mig.Down(m.steps)
}
