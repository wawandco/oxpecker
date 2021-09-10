package test

import (
	"context"

	"github.com/wawandco/ox/plugins"
)

// AfterTester is suited for things that need to run after the tests
// cleanup and organization things, maybe reporting or collecting metrics.
type AfterTester interface {
	plugins.Plugin
	RunAfterTest(context.Context, string, []string) error
}
