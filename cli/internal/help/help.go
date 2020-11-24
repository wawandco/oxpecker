package help

import (
	"context"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

// Help command that prints
type Help struct {
	commands []plugins.Plugin
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
	command := h.findCommand(args)
	if command == nil {
		h.printTopLevel()
		return nil
	}

	h.printSingle(command)
	return nil
}

func (h *Help) findCommand(args []string) plugins.Plugin {
	if len(args) < 2 {
		return nil
	}

	var command plugins.Plugin
	name := args[1]
	for _, c := range h.commands {
		if c.Name() != name {
			continue
		}

		command = c
		break
	}

	return command
}

// Receives the plugins and stores the Commands for
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
