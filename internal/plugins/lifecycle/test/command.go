// test package contains the tooling for the test
// command on the cli. The goal of this package is to provide the
// structure for test commands to run and be organized.
package test

import (
	"context"
	"fmt"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

var _ plugins.Plugin = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)

type Command struct {
	beforeTesters []BeforeTester
	testers       []Tester
	afterTesters  []AfterTester
}

func (c Command) Name() string {
	return "test"
}

func (c Command) HelpText() string {
	return "provides the structure for test commands to run and be organized"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	var err error
	for _, bt := range c.beforeTesters {
		err = bt.RunBeforeTest(ctx, root, args[1:])
		if err != nil {
			fmt.Printf("Error running %v before tester: %v\n", bt.Name(), err)
			break
		}
	}

	if err == nil {
		for _, tt := range c.testers {
			err = tt.Test(ctx, root, args[1:])
			if err != nil {
				break
			}
		}
	}

	for _, at := range c.afterTesters {
		err := at.RunAfterTest(ctx, root, args[1:])
		if err != nil {
			fmt.Printf("error running %v after tester: %v\n", at.Name(), err)
		}
	}

	return nil
}
