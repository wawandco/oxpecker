package yarn

import "github.com/paganotoni/x/internal/plugins"

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) Name() string {
	return "pop"
}
