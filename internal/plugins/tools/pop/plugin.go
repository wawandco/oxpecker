package pop

import (
	"github.com/wawandco/oxpecker/plugins"
)

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct {
	// subcommands we will invoke depending on parameters
	// these are filled when Receive is called.
	subcommands []plugins.Command
}

func (p *Plugin) Name() string {
	return "pop"
}

func (p *Plugin) ParentName() string {
	return ""
}
