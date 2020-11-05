package cli

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/commands/build"
	"github.com/paganotoni/x/commands/dev"
	"github.com/paganotoni/x/tools/compiler"
	"github.com/paganotoni/x/tools/packr"
	"github.com/paganotoni/x/tools/refresh"
	"github.com/paganotoni/x/tools/webpack"
)

// cli is the CLI wrapper for our tool. It is in charge
// for articulating different commands, finding it and also
// taking care of the CLI iteraction.
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
		fmt.Println("no command provided, please provide one")
		return nil
	}

	command := c.findCommand(args[1])
	if command == nil {
		fmt.Printf("did not find %s command\n", args[1])
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
		refresh.Tool{},
		packr.Tool{},
		compiler.Tool{},
	}

	return cli{
		commands: []command{
			build.New(tools),
			dev.New(tools),
		},
	}
}
