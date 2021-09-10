package db

import (
	"context"
	"errors"

	"github.com/gobuffalo/pop/v5"
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/log"
)

type ResetCommand struct {
	connectionName string

	flags *pflag.FlagSet
}

func (d ResetCommand) Name() string {
	return "reset"
}

func (d ResetCommand) HelpText() string {
	return "resets database specified in GO_ENV or --conn"
}

func (d ResetCommand) ParentName() string {
	return "db"
}

func (d *ResetCommand) Run(ctx context.Context, root string, args []string) error {
	conn := pop.Connections[d.connectionName]
	if conn == nil {
		return ErrConnectionNotFound
	}

	resetter := conn.Dialect
	if resetter == nil {
		return errors.New("provided connection is not a Resetter")
	}

	err := resetter.DropDB()
	if err != nil {
		log.Warnf("could not drop database: %v\n", err)
	}

	return resetter.CreateDB()
}

// RunBeforeTests will be invoked to reset the test database before
// tests run.
func (d *ResetCommand) RunBeforeTest(ctx context.Context, root string, args []string) error {
	err := pop.LoadConfigFile()
	if err != nil {
		return err
	}

	conn := pop.Connections["test"]
	if conn == nil {
		return ErrConnectionNotFound
	}

	resetter := conn.Dialect
	if resetter == nil {
		return errors.New("provided connection is not a Resetter")
	}

	err = resetter.DropDB()
	if err != nil {
		log.Warnf("could not drop database: %v\n", err)
	}

	return resetter.CreateDB()
}

func (d *ResetCommand) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.StringVarP(&d.connectionName, "conn", "", "development", "the name of the connection to use")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *ResetCommand) Flags() *pflag.FlagSet {
	return d.flags
}
