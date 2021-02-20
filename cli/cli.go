package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/info"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxpecker/tools/cli/help"
)

// cli is the CLI wrapper for our tool. It is in charge
// for articulating different commands, finding it and also
// taking care of the CLI iteraction.
type cli struct {
	root    string
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

		if command.Name() != name {
			continue
		}

		return command
	}

	return nil
}

// Runs the CLI or cmd/ox/main.go
func (c *cli) Wrap(ctx context.Context, pwd string, args []string) error {
	// Not sure if we should do this here or somewhere
	// else, these are some environment variables to be set
	// and other things to check.
	os.Setenv("GO111MODULE", "on") // Modules must be ON
	os.Setenv("CGO_ENABLED", "0")  // CGO disabled

	name, err := info.ModuleName()
	if err != nil {
		fmt.Printf("[info] could not determine module name: %v\n", err)
	}

	exception := name == "github.com/wawandco/oxpecker" || name == ""

	path := filepath.Join("cmd", "ox", "main.go")
	_, err = os.Stat(path)
	if err != nil || exception {
		fmt.Print("[info] Using wawandco/oxpecker/cmd/ox \n\n")
		return c.Run(ctx, c.root, args)
	}

	bargs := []string{"run", path}
	bargs = append(bargs, args[1:]...)

	cmd := exec.CommandContext(ctx, "go", bargs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (c *cli) Run(ctx context.Context, pwd string, args []string) error {
	if len(args) < 2 {
		fmt.Println("no command provided, please provide one")
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
		fmt.Printf("did not find %s command\n", args[1])
		return nil
	}

	return command.Run(ctx, c.root, args[1:])
}

// New creates a CLI with the passed root and plugins. This becomes handy
// when specifying your own plugins.
func New() *cli {
	c := &cli{
		Plugins: []plugins.Plugin{
			help.Command{},
		},
	}

	return c
}
