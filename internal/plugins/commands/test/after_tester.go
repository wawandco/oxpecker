package test

import "context"

// AfterTester is suited for things that need to run after the tests
// cleanup and organization things, maybe reporting or collecting metrics.
type AfterTester interface {
	RunAfterTest(context.Context, string, []string) error
}
