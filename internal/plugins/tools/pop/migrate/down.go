package migrate

import (
	"context"
	"os"

	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type MigrateDown struct {
	configFile     string
	connectionName string
	migrationsPath string
	steps          int
	flags          *pflag.FlagSet
}

func (mu MigrateDown) Name() string {
	return "migrate/down"
}

func (mu MigrateDown) CommandName() string {
	return "down"
}

// Run will run migrations on the current folder, it will look for the
// migrations folder and attempt to run the migrations using internal
// pop tooling
func (mu *MigrateDown) RunMigrations(ctx context.Context, root string, args []string) error {
	cb, err := os.OpenFile(mu.configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	err = pop.LoadFrom(cb)
	if err != nil {
		return err
	}

	conn := pop.Connections[mu.connectionName]
	if conn == nil {
		return ErrCouldNotFindConnection
	}

	mig, err := pop.NewFileMigrator(mu.migrationsPath, conn)
	if err != nil {
		return err
	}

	// Should be down
	return mig.Down(mu.steps)
}

func (mu *MigrateDown) ParseFlags(args []string) {
	mu.flags = pflag.NewFlagSet(mu.Name(), pflag.ContinueOnError)

	mu.flags.StringVarP(&mu.connectionName, "conn", "", "development", "the name of the connection to use")
	mu.flags.StringVarP(&mu.migrationsPath, "folder", "", "migrations", "the path to the migrations")
	mu.flags.StringVarP(&mu.configFile, "config", "", "config/database.yml", "direction to run the migrations to")
	mu.flags.IntVarP(&mu.steps, "steps", "s", 1, "how many migrations to run")
	mu.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (mu *MigrateDown) Flags() *pflag.FlagSet {
	return mu.flags
}
