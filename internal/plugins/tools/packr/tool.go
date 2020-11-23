package packr

import "github.com/paganotoni/oxpecker/internal/plugins"

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

// Plugin provides operations with Plugin for the CLI.
type Plugin struct{}

func (w Plugin) Name() string {
	return "Packr"
}
