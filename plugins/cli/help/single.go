package help

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/paganotoni/oxpecker/plugins"
)

// printSingle prints help details for a passed plugin
// Usage, Subcommands and Flags.
func (h *Help) printSingle(command plugins.Plugin, names []string) {

	if th, ok := command.(plugins.HelpTexter); ok {
		fmt.Printf("%v\n\n", th.HelpText())
	}

	fmt.Println("Usage:")
	usage := fmt.Sprintf("  ox %v \n", command.Name())
	th, isSubcommander := command.(plugins.Subcommander)

	_, isSubcommand := command.(plugins.Subcommand)
	if isSubcommand {
		usage = fmt.Sprintf("  ox %v \n", strings.Join(names, " "))
	}

	if isSubcommander {
		usage = fmt.Sprintf("  ox %v [subcommand]\n", command.Name())
	}

	fmt.Println(usage)

	if isSubcommander {
		w := new(tabwriter.Writer)
		defer w.Flush()

		w.Init(os.Stdout, 8, 8, 3, '\t', 0)
		fmt.Println("Subcommands:")

		for _, scomm := range th.Subcommands() {
			sc, ok := scomm.(plugins.Subcommand)
			if !ok {
				continue
			}

			helpText := ""
			if ht, ok := scomm.(plugins.HelpTexter); ok {
				helpText = ht.HelpText()
			}

			fmt.Fprintf(w, "  %v\t%v\n", sc.SubcommandName(), helpText)
		}
	}

	if th, ok := command.(plugins.FlagParser); ok {
		fmt.Println("Flags:")

		flags := th.Flags()
		flags.SetOutput(os.Stderr)
		flags.PrintDefaults()
		fmt.Println("")

		return
	}

}
