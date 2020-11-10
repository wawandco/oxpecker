package test

import "context"

// BeforeTester interface is suited for those tasks that need to happen
// before the tests run, things like setting up environment variables,
// clearing the database or other cleanup tasks.
type BeforeTester interface {
	RunBeforeTest(context.Context, string, []string) error
}
