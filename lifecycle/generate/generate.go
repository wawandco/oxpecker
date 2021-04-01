// generate package provides the base of the generate command
// which allows to run generators for tools.
package generate

import (
	"context"
	"errors"

	"github.com/wawandco/oxpecker/internal/log"
	"github.com/wawandco/oxpecker/plugins"
)

//HelpText returns the help Text of build function

var _ plugins.Command = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)

type Command struct {
	generators []Generator
}

func (c Command) Name() string {
	return "generate"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "Allows to invoke registered generator plugins"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		log.Error("no generator name specified")
		return nil
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
		return errors.New("generator not found")
	}

	return generator.Generate(ctx, root, args)
}

func (c *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {
		ptool, ok := plugin.(Generator)
		if !ok {
			continue
		}

		c.generators = append(c.generators, ptool)
	}
}
