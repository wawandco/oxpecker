package help

import (
	"context"
	"errors"
	"fmt"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

var ErrSubCommandNotFound = errors.New("Subcommand not found")

// Help command that prints
type Help struct {
	commands []plugins.Command
}

func (h Help) Name() string {
	return "help"
}

// HelpText for the Help command
func (h Help) HelpText() string {
	return "prints help text for the commands registered"
}

// Run the help command
func (h *Help) Run(ctx context.Context, root string, args []string) error {
	command, names := h.findCommand(args)
	if command == nil {
		h.printTopLevel()
		return nil
	}
	fmt.Println(names)
	h.printSingle(command, names)

	return nil
}

func (h *Help) findCommand(args []string) (plugins.Plugin, []string) {
	if len(args) < 2 {
		return nil, nil
	}

	var commands = h.commands
	var command plugins.Plugin
	var argIndex = 1
	var fndNames []string

	for {
		var name = args[argIndex]
		for _, c := range commands {
			// TODO: If its a subcommand check also the SubcommandName
			if c.Name() != name {
				continue
			}
			fndNames = append(fndNames, c.Name())
			command = c
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

		var sbcm []plugins.Command
		for _, subc := range sc.Subcommands() {
			sbcm = append(sbcm, subc.(plugins.Command))
		}

		commands = sbcm
	}

	return command, fndNames
}

// Receive the plugins and stores the Commands for
// later usage on the help text.
func (h *Help) Receive(pl []plugins.Plugin) {
	for _, plugin := range pl {

		if _, ok := plugin.(plugins.Subcommand); ok {
			continue
		}

		if ht, ok := plugin.(plugins.Command); ok {
			h.commands = append(h.commands, ht)
		}
	}
}
