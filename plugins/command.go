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

// https://github.com/gobuffalo/buffalo-cli/blob/81f172713e1182412f27a0b128160386e04cd39b/internal/garlic/run.go#L28
