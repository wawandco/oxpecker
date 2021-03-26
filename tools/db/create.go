package db

import (
	"context"

	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type CreateCommand struct {
	connectionName string
	flags          *pflag.FlagSet
}

func (d CreateCommand) Name() string {
	return "create"
}

func (d CreateCommand) HelpText() string {
	return "creates database in GO_ENV or --conn flag"
}

func (d CreateCommand) ParentName() string {
	return "db"
}

func (d *CreateCommand) Run(ctx context.Context, root string, args []string) error {
	conn := pop.Connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	return conn.Dialect.CreateDB()
}

func (d *CreateCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *CreateCommand) Flags() *pflag.FlagSet {
	return d.flags
}
