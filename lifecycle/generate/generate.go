// generate package provides the base of the generate command
// which allows to run generators for tools.
package generate

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins"
)

//HelpText returns the help Text of build function

var _ plugins.Command = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)

type Command struct {
	generators      []Generator
	aftergenerators []AfterGenerator
}

func (c Command) Name() string {
	return "generate"
}

func (c Command) Alias() string {
	return "g"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "Allows to invoke registered generator plugins"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		c.list()

		return fmt.Errorf("no generator name specified")
	}

	name := args[1]
	var generator Generator
	//Run each of the fixers registered.
	for _, gen := range c.generators {
		if gen.InvocationName() != name {
			continue
		}

		generator = gen
		break
	}

	if generator == nil {
		c.list()
		return fmt.Errorf("generator `%v` not found", name)
	}

	err := generator.Generate(ctx, root, args)
	for _, ag := range c.aftergenerators {
		aerr := ag.AfterGenerate(ctx, root, args)
		if aerr != nil {
			log.Errorf("Error running after generator %v: %v", ag.Name(), aerr)
		}
	}

	return err
}

func (c Command) list() {
	w := new(tabwriter.Writer)
	defer w.Flush()

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 3, '\t', 0)
	fmt.Printf("Available Generators:\n\n")
	fmt.Fprintf(w, "  Name\tPlugin\n")
	fmt.Fprintf(w, "  ----\t------\n")
	for _, plugin := range c.generators {
		helpText := ""
		if ht, ok := plugin.(plugins.HelpTexter); ok {
			helpText = ht.HelpText()
		}

		fmt.Fprintf(w, "  %v\t%v\t%v\n", plugin.InvocationName(), plugin.Name(), helpText)
	}

	fmt.Fprintf(w, "\n")
}

func (c *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		ptool, ok := plugin.(Generator)
		if !ok {
			continue
		}

		c.generators = append(c.generators, ptool)
	}

	for _, plugin := range plugins {
		ptool, ok := plugin.(AfterGenerator)
		if !ok {
			continue
		}

		c.aftergenerators = append(c.aftergenerators, ptool)
	}
}
