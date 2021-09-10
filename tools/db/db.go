// db packs all db operations under this top level command.
package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gobuffalo/pop/v5"
	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)
var _ plugins.Subcommander = (*Command)(nil)

var ErrConnectionNotFound = errors.New("connection not found")

type Command struct {
	subcommands []plugins.Command
}

func (c Command) Name() string {
	return "db"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "database operation commands"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		log.Error("no subcommand specified, please use `db [subcommand]` to run one of the db subcommands.")
		return nil
	}

	err := pop.LoadConfigFile()
	if err != nil {
		log.Error(err.Error())
	}

	name := args[1]
	var subcommand plugins.Command
	for _, sub := range c.subcommands {
		if sub.Name() != name {
			continue
		}

		subcommand = sub
		break
	}

	if subcommand == nil {
		return fmt.Errorf("subcommand `%v` not found", name)
	}

	return subcommand.Run(ctx, root, args)
}

func (c *Command) Receive(pls []plugins.Plugin) {
	for _, plugin := range pls {
		ptool, ok := plugin.(plugins.Command)
		if !ok || ptool.ParentName() != c.Name() {
			continue
		}

		c.subcommands = append(c.subcommands, ptool)
	}
}

func (c *Command) Subcommands() []plugins.Command {
	return c.subcommands
}

func (c *Command) FindRoot() string {
	root := info.RootFolder()
	if root != "" {
		return root
	}

	root, err := os.Getwd()
	if err != nil {
		return ""
	}

	return root
}

func Plugins() []plugins.Plugin {
	var result []plugins.Plugin

	result = append(result, &Command{})
	result = append(result, &CreateCommand{})
	result = append(result, &DropCommand{})
	result = append(result, &ResetCommand{})

	return result
}
