// Standard package provides plugin that performs operations
// of the Go standard toolset. Things like compiling or running
// the base `go` prefixed commands.
package standard

import (
	"github.com/paganotoni/oxpecker/internal/plugins"
	"github.com/spf13/pflag"
)

var (
	// These are the interfaces we know that this
	// plugin must satisfy for its correct functionality
	_ plugins.Plugin     = (*Plugin)(nil)
	_ plugins.FlagParser = (*Plugin)(nil)
)

// Compiler takes care of compiling the go binary, this might be one of
// the last steps when we have done other phases of the build process.
type Plugin struct {
	output string
	flags  *pflag.FlagSet
}

func (g Plugin) Name() string {
	return "standard"
}
