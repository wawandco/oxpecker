// fix package contains the logics of the fix operations, fix operations
// are in charge of adapting our source code to comply with newer versions
// of the CLI.
package fix

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

// Things to Fix:

// 1. models/models.go has changed its structure not to use an init function to
// set the database, it now provides a method to return the database connection

type Command struct {
	fixers []Fixer
}

func (c Command) Name() string {
	return "fix"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Println("Running [fix] command")
	//Run each of the fixers registered.
	for _, fixer := range c.fixers {
		fmt.Printf("Fixer: %v\n", fixer.Name())
	}

	return nil
}

func (c *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {

		if ptool, ok := plugin.(Fixer); ok {
			b.fixers = append(b.fixers, ptool)
		}

	}
}
