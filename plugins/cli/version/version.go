package version

import (
	"context"
	"fmt"

	"github.com/paganotoni/oxpecker/plugins"
)

var (
	// The version of the CLI
	version = "0.0.1"
)

var (
	_ plugins.Command = (*Version)(nil)
)

// Version command will print X version.
type Version struct{}

func (b Version) Name() string {
	return "version"
}
func (b Version) HelpText() string {
	return "retuns the curren version of Oxpecker CLI"
}

// Run prints the version of the Oxpecker cli by using the
// Versioner in the command, one for the x tool.
func (b *Version) Run(ctx context.Context, root string, args []string) error {
	fmt.Printf("Oxpecker version %v\n", version)

	return nil
}
