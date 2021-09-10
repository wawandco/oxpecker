package help

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/wawandco/ox/plugins"
)

// printSingle prints help details for a passed plugin
// Usage, Subcommands and Flags.
func (h *Command) printSingle(command plugins.Command, names []string) {

	if th, ok := command.(plugins.HelpTexter); ok {
		fmt.Printf("%v\n\n", th.HelpText())
	}

	fmt.Println("Usage:")
	usage := fmt.Sprintf("  ox %v \n", command.Name())

	if command.ParentName() != "" {
		usage = fmt.Sprintf("  ox %v \n", strings.Join(names, " "))
	}

	th, isSubcommander := command.(plugins.Subcommander)
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
			if scomm.ParentName() == "" {
				continue
			}

			helpText := ""
			if ht, ok := scomm.(plugins.HelpTexter); ok {
				helpText = ht.HelpText()
			}

			fmt.Fprintf(w, "  %v\t%v\n", scomm.Name(), helpText)
		}
	}
	if th, ok := command.(plugins.Aliaser); ok {
		fmt.Println("Alias:")
		fmt.Println(th.Alias())
		fmt.Println("")
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
