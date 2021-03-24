package db

import (
	"context"

	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type DropCommand struct {
	connectionName string

	flags *pflag.FlagSet
}

func (d DropCommand) Name() string {
	return "drop"
}

func (d DropCommand) HelpText() string {
	return "drops database in GO_ENV or --conn flag"
}

func (d DropCommand) ParentName() string {
	return "db"
}

func (d *DropCommand) Run(ctx context.Context, root string, args []string) error {
	conn := pop.Connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	return conn.Dialect.DropDB()
}

func (d *DropCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *DropCommand) Flags() *pflag.FlagSet {
	return d.flags
}
