package yarn

import "github.com/wawandco/oxpecker/plugins"

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) Name() string {
	return "yarn"
}
