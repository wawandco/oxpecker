package dev

import (
	"context"

	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins"
	"golang.org/x/sync/errgroup"
)

var _ plugins.Command = (*Command)(nil)

// Command is the dev command, it runs the dev plugins, each one on a different
// go routine. the detail to what happen on each of these plugins is up to
// each of the Developer plugins.
type Command struct {
	developers []Developer
	beforeDevs []BeforeDeveloper
}

func (d Command) Name() string {
	return "dev"
}

func (c Command) Alias() string {
	return "d"
}

func (d Command) ParentName() string {
	return ""
}

//HelpText returns the help Text of build function
func (d Command) HelpText() string {
	return "calls NPM or yarn to start webpack watching the assetst"
}

// Run calls each of the beforedeveloper plugins and then
// executes Developer plugins in parallel.
func (d *Command) Run(ctx context.Context, root string, args []string) error {
	for _, bd := range d.beforeDevs {
		err := bd.BeforeDevelop(ctx, root)
		if err != nil {
			return err
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &errgroup.Group{}
	for _, d := range d.developers {
		g := d.Develop
		wg.Go(func() error {
			err := g(ctx, root)
			if err != nil {
				log.Error(err.Error())
			}

			return nil
		})
	}

	return wg.Wait()
}

// Receive Developer and BeforeDeveloper plugins and store these
// in the Command to be used when the command is invoked.
func (d *Command) Receive(plugins []plugins.Plugin) {
	for _, tool := range plugins {
		if ptool, ok := tool.(Developer); ok {
			d.developers = append(d.developers, ptool)
		}

		if bdev, ok := tool.(BeforeDeveloper); ok {
			d.beforeDevs = append(d.beforeDevs, bdev)
		}
	}
}
