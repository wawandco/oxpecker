package x

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

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
	fmt.Printf("Oxpecker allows to build apps with ease\n\n")
	fmt.Println("Usage:")
	fmt.Printf("  ox [command]\n\n")

	w := new(tabwriter.Writer)
	defer w.Flush()

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 3, '\t', 0)
	fmt.Println("Commands:")

	for _, plugin := range h.commands {
		helpText := ""
		if ht, ok := plugin.(plugins.HelpTexter); ok {
			helpText = ht.HelpText()
		}

		fmt.Fprintf(w, "  %v\t%v\n", plugin.Name(), helpText)
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
