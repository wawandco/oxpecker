package plugins

import (
	"context"
)

// Command interface for commands that the CLI provides.
// a command is one of the top cli elements (build, fix, generate ...)
type Command interface {
	Plugin

	// Run the command with the passed context, root and args.
	Run(context.Context, string, []string) error
}

// Command namer is an interface
type CommandNamer interface {
	Command

	// Command name returns a specific name for the plugin to be
	// used to identify the command.
	CommandName() string
}

// SubcommandNamer allows to identify those commands that will not be added by
// the CLI as top level commands but rather those top level commands will organize
// as subcommands.
type Subcommand interface {
	Command
	// SubcommandName is the name for the
	SubcommandName() string

	// Run the command with the passed context, root and args.
	Run(context.Context, string, []string) error
}

// Subcommander allows a plugin to say which are its subcommands.
type Subcommander interface {
	Command

	Subcommands() []Subcommand
}
