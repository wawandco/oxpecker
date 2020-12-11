package migrate

import (
	"context"
	"errors"

	"github.com/wawandco/oxpecker/plugins"
)

var (
	_ plugins.Command = (*Plugin)(nil)

	ErrCouldNotFindConnection = errors.New("could not find connection by name")
	ErrNotEnoughArgs          = errors.New("not enough args, please specify direction p.e ox pop migrate up")
	ErrInvalidInstruction     = errors.New("invalid instruction for migrate command")
)

type Plugin struct {
	migrators []Migrator
}

//HelpText resturns the help Text of build function
func (m Plugin) HelpText() string {
	return "Runs migrations on the current folder"
}

func (m *Plugin) Name() string {
	return "migrate"
}

func (m *Plugin) ParentName() string {
	return "pop"
}

func (m *Plugin) SubcommandName() string {
	return "migrate"
}

func (m *Plugin) Run(ctx context.Context, root string, args []string) error {

	if len(args) < 2 {
		return ErrNotEnoughArgs
	}

	name := args[1]
	for _, mig := range m.migrators {
		if mig.CommandName() != name {
			continue
		}

		return mig.RunMigrations(ctx, root, args)
	}

	return ErrInvalidInstruction
}

func (m *Plugin) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		pl, ok := plugin.(Migrator)
		if !ok {
			continue
		}

		m.migrators = append(m.migrators, pl)
	}
}
