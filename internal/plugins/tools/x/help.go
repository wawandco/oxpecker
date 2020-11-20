package x

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

// Help command that prints
type Help struct {
	commands []plugins.Plugin
}

func (h Help) Name() string {
	return "help"
}

func (h Help) HelpText() string {
	return "prints commands help text"
}

// Run the help command
func (h *Help) Run(ctx context.Context, root string, args []string) error {
	fmt.Printf("X allows to build apps with ease\n\n")
	fmt.Println("Usage:")
	fmt.Printf("  x [command]\n\n")

	fmt.Printf("Commands:\n")
	for _, plugin := range h.commands {
		helpText := ""
		if ht, ok := plugin.(plugins.HelpTexter); ok {
			helpText = ht.HelpText()
		}

		fmt.Printf("  %v - %v\n", plugin.Name(), helpText)
	}

	return nil
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
