package cli

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/plugins"
	"github.com/paganotoni/x/plugins/build"
	"github.com/paganotoni/x/plugins/compiler"
	"github.com/paganotoni/x/plugins/dev"
	"github.com/paganotoni/x/plugins/packr"
	"github.com/paganotoni/x/plugins/refresh"
	"github.com/paganotoni/x/plugins/webpack"
)

// defaultPlugins is the list of default plugins that will
// be used by default.
var defaultPlugins = []plugins.Plugin{

	//IMPORTANT: order matters!
	webpack.Tool{},
	refresh.Tool{},
	packr.Tool{},
	compiler.Tool{},

	// Commands are plugins
	&build.Command{},
	&dev.Command{},
}

// cli is the CLI wrapper for our tool. It is in charge
// for articulating different commands, finding it and also
// taking care of the CLI iteraction.
type cli struct {
	root    string
	plugins []plugins.Plugin
}

// findCommand looks in the plugins for a command
// with the passed name.
func (c cli) findCommand(name string) plugins.Command {
	for _, cm := range c.plugins {
		command, ok := cm.(plugins.Command)
		if !ok {
			continue
		}

		if command.Name() != name {
			continue
		}

		return command
	}

	return nil
}

// Runs the CLI
func (c cli) Run(args []string) error {
	if len(args) < 2 {
		fmt.Println("no command provided, please provide one")
		return nil
	}

	command := c.findCommand(args[1])
	if command == nil {
		fmt.Printf("did not find %s command\n", args[1])
		return nil
	}

	if pr, ok := command.(plugins.PluginReceiver); ok {
		pr.Receive(c.plugins)
	}

	ctx := context.Background()
	return command.Run(ctx, c.root, args[1:])
}

// New creates a CLI with the passed root and plugins. This becomes handy
// when specifying your own plugins.
func New(root string, plugins []plugins.Plugin) cli {
	return cli{
		root:    root,
		plugins: plugins,
	}
}

// NewWithRoot creates a CLI with the root passed and
// default set of plugins.
func NewWithRoot(root string) cli {
	return cli{
		root:    root,
		plugins: defaultPlugins,
	}
}
