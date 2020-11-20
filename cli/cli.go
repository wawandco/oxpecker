package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/paganotoni/x/internal/plugins"
	"github.com/paganotoni/x/internal/plugins/lifecycle/build"
	"github.com/paganotoni/x/internal/plugins/lifecycle/dev"
	"github.com/paganotoni/x/internal/plugins/lifecycle/fix"
	"github.com/paganotoni/x/internal/plugins/lifecycle/test"
	"github.com/paganotoni/x/internal/plugins/lifecycle/version"
	"github.com/paganotoni/x/internal/plugins/tools/packr"
	"github.com/paganotoni/x/internal/plugins/tools/pop"
	"github.com/paganotoni/x/internal/plugins/tools/pop/migrate"
	"github.com/paganotoni/x/internal/plugins/tools/refresh"
	"github.com/paganotoni/x/internal/plugins/tools/standard"
	"github.com/paganotoni/x/internal/plugins/tools/webpack"
	"github.com/paganotoni/x/internal/plugins/tools/x"
	"github.com/paganotoni/x/internal/plugins/tools/yarn"
)

// defaultPlugins is the list of default plugins that will
// be used by default.
var defaultPlugins = []plugins.Plugin{
	// IMPORTANT: order matters!
	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&packr.Plugin{},
	&pop.Plugin{},
	&migrate.Plugin{},
	&standard.Plugin{},
	&yarn.Plugin{},
	&x.Fixer{},

	// Developer Lifecycle plugins
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
	&version.Command{Version: "1.1.1"},
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
func (c *cli) findCommand(name string) plugins.Command {
	for _, cm := range c.plugins {
		// We skip subcommands on this case
		// those will be wired by the parent command implementing
		// Receive.
		if _, ok := cm.(plugins.Subcommand); ok {
			continue
		}

		command, ok := cm.(plugins.Command)
		if !ok {
			continue
		}

		pluginName := command.Name()
		if pn, ok := cm.(plugins.CommandNamer); ok {
			pluginName = pn.CommandName()
		}

		if pluginName != name {
			continue
		}

		return command
	}

	return nil
}

// Runs the CLI
func (c *cli) Run(args []string) error {

	// IMPORTANT: Incorporate the plugin system by taking a look at this.
	// https://github.com/gobuffalo/buffalo-cli/blob/81f172713e1182412f27a0b128160386e04cd39b/internal/garlic/run.go#L28

	// Not sure if we should do this here or somewhere
	// else, these are some environment variables to be set
	// and other things to check.
	os.Setenv("GO111MODULE", "on") // Modules must be ON
	os.Setenv("CGO_ENABLED", "0")  // CGO disabled

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

	if pf, ok := command.(plugins.FlagParser); ok {
		err := pf.ParseFlags(args[1:])
		if err != nil {
			fmt.Println(err)
		}
	}

	ctx := context.Background()
	return command.Run(ctx, c.root, args[1:])
}

// New creates a CLI with the passed root and plugins. This becomes handy
// when specifying your own plugins.
func New(root string, plugins []plugins.Plugin) *cli {
	return &cli{
		root:    root,
		plugins: plugins,
	}
}

// NewWithRoot creates a CLI with the root passed and
// default set of plugins.
func NewWithRoot(root string) *cli {
	return &cli{
		root:    root,
		plugins: defaultPlugins,
	}
}
