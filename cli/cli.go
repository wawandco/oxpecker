package cli

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins"
	"github.com/wawandco/ox/tools/cli/help"
)

// cli is the CLI wrapper for our tool. It is in charge
// for articulating different commands, finding it and also
// taking care of the CLI iteraction.
type cli struct {
	Plugins []plugins.Plugin
}

// findCommand looks in the plugins for a command
// with the passed name.
func (c *cli) findCommand(name string) plugins.Command {
	for _, cm := range c.Plugins {
		// We skip subcommands on this case
		// those will be wired by the parent command implementing
		// Receive.
		command, ok := cm.(plugins.Command)
		if !ok || command.ParentName() != "" {
			continue
		}

		alias, ok := cm.(plugins.Aliaser)
		if ok && alias.Alias() == name {
			return command
		}

		if command.Name() == name {
			return command
		}

	}

	return nil
}

// Runs the CLI or cmd/ox/main.go
func (c *cli) Wrap(ctx context.Context, args []string) error {
	// Not sure if we should do this here or somewhere
	// else, these are some environment variables to be set
	// and other things to check.
	os.Setenv("GO111MODULE", "on") // Modules must be ON
	os.Setenv("CGO_ENABLED", "0")  // CGO disabled

	path := filepath.Join("cmd", "ox", "main.go")
	_, err := os.Stat(path)
	name := info.ModuleName()
	if err != nil || name == "" || name == "github.com/wawandco/ox" {
		log.Info("Using github.com/wawandco/ox/cmd/ox \n")
		return c.Run(ctx, args)
	}

	bargs := []string{"run", path}
	bargs = append(bargs, args[1:]...)

	log.Infof("Using %v \n", path)
	cmd := exec.CommandContext(ctx, "go", bargs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (c *cli) Run(ctx context.Context, args []string) error {
	if len(args) < 2 {
		log.Error("no command provided, please provide one")
		return nil
	}

	// Passing args and plugins to those plugins that require them
	for _, plugin := range c.Plugins {
		pf, ok := plugin.(plugins.FlagParser)
		if ok {
			pf.ParseFlags(args[1:])
		}
		pr, ok := plugin.(plugins.PluginReceiver)
		if ok {
			pr.Receive(c.Plugins)
		}
	}

	command := c.findCommand(args[1])
	if command == nil {
		// TODO: print help ?
		log.Infof("did not find %s command\n", args[1])
		return nil
	}

	// Commands that require running within the ox directory
	// may require its root to be determined with the go.mod. However
	// some other commands may want to determine the root by themself,
	// doing os.Getwd or something similar. The latter ones are RootFinders.
	root := info.RootFolder()
	rf, ok := command.(plugins.RootFinder)
	if root == "" && !ok {
		return errors.New("go.mod not found")
	}

	if root == "" {
		root = rf.FindRoot()
	}

	return command.Run(ctx, root, args[1:])
}

// Use passed Pugins by appending these to the
// plugins list inside the CLI.
func (cl *cli) Use(plugins ...plugins.Plugin) {
	cl.Plugins = append(cl.Plugins, plugins...)
}

// Remove looks in the plugins list and removes plugins that
// match passed names.
func (cl *cli) Remove(names ...string) {
	result := []plugins.Plugin{}
	for _, pl := range cl.Plugins {
		var found bool
		for _, restricted := range names {
			if pl.Name() == restricted {
				found = true
			}
		}

		if found {
			continue
		}

		result = append(result, pl)
	}

	cl.Plugins = result
}

// Clear the plugin list of the CLI.
func (cl *cli) Clear() {
	cl.Plugins = []plugins.Plugin{}
}

// New creates a CLI with the passed root and plugins. This becomes handy
// when specifying your own plugins.
func New() *cli {
	log.Warn("cli.New() is deprecated, se with caution.")

	return &cli{
		Plugins: []plugins.Plugin{
			help.Command{},
		},
	}
}
