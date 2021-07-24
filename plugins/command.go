package plugins

import (
	"context"
)

// Command interface for commands that the CLI provides.
// a command is one of the top cli elements (build, fix, generate ...)
type Command interface {
	Plugin

	// ParentName allows to identify subcommands and its parents.
	ParentName() string

	// Run the command with the passed context, root and args.
	Run(context.Context, string, []string) error
}

// RootFinder allows some commands not to depend on the go.mod to determine the root folder,
// this comes handy for commands like New and Version.
type RootFinder interface {
	Plugin

	// FindRoot returns the path to consider as root.
	FindRoot() string
}

// Aiaser allows commands to have aliases
type Aliaser interface {
	Alias() string
}
