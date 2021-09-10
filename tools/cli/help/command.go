package help

import (
	"context"
	"errors"

	"github.com/wawandco/ox/plugins"
)

var (
	// Help is a Command
	_ plugins.Command = (*Command)(nil)

	ErrSubCommandNotFound = errors.New("subcommand not found")
)

// Help command that prints
type Command struct {
	commands []plugins.Command
}

func (h Command) Name() string {
	return "help"
}

func (c Command) Alias() string {
	return "h"
}

func (h Command) ParentName() string {
	return ""
}

// HelpText for the Help command
func (h Command) HelpText() string {
	return "prints help text for the commands registered"
}

// Run the help command
func (h *Command) Run(ctx context.Context, root string, args []string) error {
	command, names := h.findCommand(args)
	if command == nil {
		h.printTopLevel()
		return nil
	}

	h.printSingle(command, names)

	return nil
}

func (h *Command) findCommand(args []string) (plugins.Command, []string) {
	if len(args) < 2 {
		return nil, nil
	}

	var commands = h.commands
	var command plugins.Command
	var argIndex = 1
	var fndNames []string

	for {
		var name = args[argIndex]
		for _, c := range commands {
			// TODO: If its a subcommand check also the SubcommandName
			a, ok := c.(plugins.Aliaser)
			if !ok {
				if c.Name() != name {
					continue
				}
				fndNames = append(fndNames, c.Name())
				command = c
				break
			}

			if a.Alias() != name && c.Name() != name {
				continue
			}

			if c.Name() == name {
				fndNames = append(fndNames, c.Name())
				command = c
				break
			}

			if a.Alias() == name {
				fndNames = append(fndNames, a.Alias())
				command = c
			}
			break
		}

		argIndex++
		if argIndex >= len(args) {
			break
		}

		sc, ok := command.(plugins.Subcommander)
		if !ok {
			break
		}

		commands = sc.Subcommands()
	}

	return command, fndNames
}

// Receive the plugins and stores the Commands for
// later usage on the help text.
func (h *Command) Receive(pl []plugins.Plugin) {
	for _, plugin := range pl {
		ht, ok := plugin.(plugins.Command)
		if ok && ht.ParentName() == "" {
			h.commands = append(h.commands, ht)
		}
	}
}
