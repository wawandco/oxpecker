package test

import "context"

// Tester runs a set of tests depending on the tools you
// want to test, this could include
// - Go test (go test ...)
// - Linting (gofmt/metalinter/milo)
// - Yarn/NPM tests
type Tester interface {
	RunTest(context.Context, string, []string) error
}
