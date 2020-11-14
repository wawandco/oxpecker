package pop

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

// Ensuring pop.Plugin is a command
var _ plugins.Command = (*Plugin)(nil)

// Run will take care of pop subcommands
// This will be invoked by the CLI with x pop [subcommand]
func (b *Plugin) Run(ctx context.Context, root string, args []string) error {
	fmt.Println("Running pop")
	return nil
}
