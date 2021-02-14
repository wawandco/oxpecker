package new

import (
	"context"
	"errors"
	"path/filepath"
	"sync"

	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)
var ErrNoNameProvided = errors.New("the name for the new app is needed")

// Command to generate New applications.
type Command struct {
	initializers      []Initializer
	afterInitializers []AfterInitializer
}

func (d Command) Name() string {
	return "new"
}

func (d Command) ParentName() string {
	return ""
}

//HelpText returns the help Text of build function
func (d Command) HelpText() string {
	return "Generates a new app with registered plugins"
}

// Run
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNoNameProvided
	}

	name := d.AppName(args)

	var dx sync.Map
	dx.Store("args", args)
	dx.Store("root", root)
	dx.Store("folder", filepath.Join(root, name))
	dx.Store("name", name)
	dx.Store("module", args[1])

	for _, ini := range d.initializers {
		err := ini.Initialize(ctx, &dx)
		if err != nil {
			return err
		}
	}

	for _, aini := range d.afterInitializers {
		err := aini.AfterInitialize(ctx, root, args)
		if err != nil {
			return err
		}
	}

	return nil
}

// Receive and store initializers
func (d *Command) Receive(plugins []plugins.Plugin) {
	for _, tool := range plugins {
		i, ok := tool.(Initializer)
		if ok {
			d.initializers = append(d.initializers, i)
		}

		ai, ok := tool.(AfterInitializer)
		if ok {
			d.afterInitializers = append(d.afterInitializers, ai)
		}
	}
}
func (d *Command) AppName(args []string) string {
	return filepath.Base(args[1])
}
