package help

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

// printTopLevel prints the top level help text with a table that contains top level
// commands (names) and descriptions.
func (h *Help) printTopLevel() {
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

}
