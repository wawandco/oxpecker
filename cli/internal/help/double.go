package help

import (
	"fmt"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

func (h *Help) printDouble(command plugins.Plugin, subcommand plugins.Subcommand) {

	if th, ok := subcommand.(plugins.HelpTexter); ok {
		fmt.Printf("%v\n\n", th.HelpText())
	}

	fmt.Println("Usage:")

	fmt.Printf("  ox %v %v", command.Name(), subcommand.SubcommandName())

}
