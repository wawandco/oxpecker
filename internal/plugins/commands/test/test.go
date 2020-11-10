// test package contains the tooling for the test
// command on the cli. The goal of this package is to provide the
// structure for test commands to run and be organized.
package test

import (
	"context"

	"github.com/paganotoni/x/internal/plugins"
)

type Command struct {
	beforeTesters []BeforeTester
	testers       []Tester
	afterTester   []AfterTester
}

func (c Command) Name() string {
	return "test"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	// Run before tester
	// Run testers
	// Run after tester
	return nil
}

func (b *Command) Receive(plugins []plugins.Plugin) {
	for _, plugin := range plugins {

		if ptool, ok := plugin.(BeforeTester); ok {
			b.beforeTesters = append(b.beforeTesters, ptool)
		}

		if ptool, ok := plugin.(Tester); ok {
			b.testers = append(b.testers, ptool)
		}

		if ptool, ok := plugin.(AfterTester); ok {
			b.afterTesters = append(b.afterTester, ptool)
		}
	}
}
