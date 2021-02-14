package db

import (
	"context"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
)

type DropCommand struct {
	connections    map[string]URLProvider
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
	conn := d.connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	if c, ok := conn.(*pop4.Connection); ok {
		return c.Dialect.DropDB()
	}

	if c, ok := conn.(*pop5.Connection); ok {
		return c.Dialect.DropDB()
	}

	return nil
}

func (d *DropCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *DropCommand) Flags() *pflag.FlagSet {
	return d.flags
}
