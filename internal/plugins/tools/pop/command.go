package pop

import (
	"context"
	"errors"

	"github.com/paganotoni/oxpecker/internal/plugins/tools/pop/migrate"
	"github.com/paganotoni/oxpecker/plugins"
)

//HelpText resturns the help Text of build function
func (b Plugin) HelpText() string {
	return "provides commands for pop common tasks"
}

// Ensuring pop.Plugin is a command
var _ plugins.Command = (*Plugin)(nil)

func (b *Plugin) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		if mig, ok := plugin.(*migrate.Plugin); ok {
			b.subcommands = append(b.subcommands, mig)
			continue
		}

		// Other subcommands
	}
}

func (b *Plugin) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return errors.New("subcommand not found")
	}

	for _, cm := range b.subcommands {
		if cm.SubcommandName() != args[1] {
			continue
		}

		if fp, ok := cm.(plugins.FlagParser); ok {
			err := fp.ParseFlags(args[1:])
			if err != nil {
				return err
			}
		}

		return cm.Run(ctx, root, args[1:])
	}

	return nil //migrate.Run(ctx, root, args[1:])
}

func (b *Plugin) Subcommands() []plugins.Subcommand {
	return b.subcommands
}
