package db

import (
	"context"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type CreateCommand struct {
	connections    map[string]URLProvider
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
	conn := d.connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	if c, ok := conn.(*pop4.Connection); ok {
		return c.Dialect.CreateDB()
	}

	if c, ok := conn.(*pop5.Connection); ok {
		return c.Dialect.CreateDB()
	}

	return nil
}

func (d *CreateCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *CreateCommand) Flags() *pflag.FlagSet {
	return d.flags
}
