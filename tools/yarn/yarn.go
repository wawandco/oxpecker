package yarn

import "github.com/wawandco/ox/plugins"

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) Name() string {
	return "yarn"
}
