package cli

import "context"

// command interface for commands that the CLI provides.
// a command is one of the top cli elements (build, fix, generate ...)
type command interface {
	Name() string
	Run(context.Context, string, []string) error
}
