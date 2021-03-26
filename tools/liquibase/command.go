package liquibase

import (
	"context"
	"errors"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)

var ErrInvalidInstruction = errors.New("Invalid instruction please specify up or down")

type Command struct {
	connectionName string
	steps          int
	connections    map[string]URLProvider
	flags          *pflag.FlagSet
}

func (lb Command) Name() string {
	return "migrate"
}

func (lb Command) ParentName() string {
	return "db"
}

func (lb Command) HelpText() string {
	return "runs Liquibase command to update database specified with --conn flag"
}

func (lb *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return lb.runUp()
	}

	direction := args[2]
	if direction == "up" {
		return lb.runUp()
	}

	if direction == "down" {
		return lb.runDown()
	}

	return ErrInvalidInstruction
}

func (lb *Command) RunBeforeTest(ctx context.Context, root string, args []string) error {
	lb.connectionName = "test"
	return lb.runUp()
}

func (lb *Command) ParseFlags(args []string) {
	lb.flags = pflag.NewFlagSet(lb.Name(), pflag.ContinueOnError)
	lb.flags.StringVarP(&lb.connectionName, "conn", "", "development", "the name of the connection to use")
	lb.flags.IntVarP(&lb.steps, "steps", "s", 0, "number of migrations to run")
	lb.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (lb *Command) Flags() *pflag.FlagSet {
	return lb.flags
}
