package new

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)
var ErrNoNameProvided = errors.New("the name for the new app is needed")

// Command to generate New applications.
type Command struct {
	// force tells whether to remove or not
	// the folder when found.
	force bool

	initializers      []Initializer
	afterInitializers []AfterInitializer

	flags *pflag.FlagSet
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

// Run each of the initializers and afterinitializers to
// compose the initial ox application.
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		return ErrNoNameProvided
	}

	name := d.AppName(args)
	folder := filepath.Join(root, name)

	if _, err := os.Stat(folder); err == nil && !d.force {
		return errors.New("folder already exist")
	}

	err := os.RemoveAll(folder)
	if err != nil {
		return err
	}

	options := Options{
		Args:   args,
		Root:   root,
		Folder: folder,
		Name:   name,
		Module: args[1],
	}

	for _, ini := range d.initializers {
		err := ini.Initialize(ctx, options)
		if err != nil {
			return err
		}
	}

	for _, aini := range d.afterInitializers {
		err := aini.AfterInitialize(ctx, options)
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

func (d *Command) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.BoolVarP(&d.force, "force", "f", false, "clear existing folder if found.")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *Command) Flags() *pflag.FlagSet {
	return d.flags
}

func (d *Command) FindRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}
