package soda

import (
	"context"
	"errors"

	"github.com/gobuffalo/packd"
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/plugins"
)

var (
	_ plugins.Command = (*Command)(nil)

	ErrCouldNotFindConnection = errors.New("could not find connection by name")
	ErrNotEnoughArgs          = errors.New("not enough args, please specify direction p.e ox pop migrate up")
	ErrInvalidInstruction     = errors.New("invalid instruction for migrate command")
)

type Command struct {
	*Logger

	steps          int
	connectionName string
	migrations     packd.Box

	flags *pflag.FlagSet
}

//HelpText returns the help Text of build function
func (m Command) HelpText() string {
	return "Uses soda to run pop migrations"
}

func (m *Command) Name() string {
	return "migrate"
}

func (m *Command) ParentName() string {
	return "db"
}

func (m *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return m.RunUp()
	}

	direction := args[2]
	if direction == "up" {
		return m.RunUp()
	}

	if direction == "down" {
		return m.RunDown()
	}

	return ErrInvalidInstruction
}

func (m *Command) ParseFlags(args []string) {
	m.flags = pflag.NewFlagSet(m.Name(), pflag.ContinueOnError)
	m.flags.StringVarP(&m.connectionName, "conn", "", "development", "the name of the connection to use")
	m.flags.IntVarP(&m.steps, "steps", "s", 0, "how many migrations to run")
	m.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (m *Command) Flags() *pflag.FlagSet {
	return m.flags
}
