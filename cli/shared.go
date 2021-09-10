package cli

import (
	"context"

	"github.com/wawandco/ox/plugins"
	"github.com/wawandco/ox/tools"
)

var (
	// shared CLI instance, its here to
	// simplify the API for custom plugins.
	shared = &cli{
		Plugins: tools.Base,
	}
)

// Use specific plugin by adding it to the
// cli plugin list.
func Use(plugins ...plugins.Plugin) {
	shared.Use(plugins...)
}

// Remove a plugin by its name form the CLI
// list of plugins.
func Remove(names ...string) {
	shared.Remove(names...)
}

// Clear all the plugins from the CLI
// useful when you just want a few plugins
// on certain environment.
func Clear() {
	shared.Clear()
}

// Run the CLI.
func Run(ctx context.Context, args []string) error {
	return shared.Run(ctx, args)
}

// Wrap runs cmd/ox/main.go if found, otherwise
// it calls cli.Run.
func Wrap(ctx context.Context, args []string) error {
	return shared.Wrap(ctx, args)
}
