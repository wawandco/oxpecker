package cli

import (
	"context"

	"github.com/paganotoni/ax/commands/build"
	"github.com/paganotoni/ax/tools/compiler"
	"github.com/paganotoni/ax/tools/packr"
	"github.com/paganotoni/ax/tools/webpack"
)

type cli struct {
	commands []command
}

func (c cli) findCommand(name string) command {
	for _, cm := range c.commands {
		if cm.Name() != name {
			continue
		}

		return cm
	}

	return nil
}

func (c cli) Run(root string, args []string) error {
	if len(args) < 2 {
		//TODO: return error
		return nil
	}

	command := c.findCommand(args[1])
	if command == nil {
		//TODO: return error
		return nil
	}

	ctx := context.Background()
	return command.Run(ctx, root, args[1:])
}

func NewCLI() cli {
	// Tools that provide capacities for the commands.
	// IMPORTANT: Order MATTERS here.
	tools := []interface{}{
		webpack.Tool{},
		packr.Tool{},
		compiler.Tool{},
	}

	return cli{
		commands: []command{
			build.New(tools),
		},
	}
}
