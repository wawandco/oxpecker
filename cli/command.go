package cli

import "context"

// command interface for commands that the CLI provides.
type command interface {
	Name() string
	Run(context.Context, string, []string) error
}
