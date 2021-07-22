package cli

import (
	"testing"

	"github.com/wawandco/oxpecker/lifecycle/generate"
	"github.com/wawandco/oxpecker/plugins"
)

func Test_CliTestingAliaser(t *testing.T) {
	plugins := []plugins.Plugin{
		&generate.Command{},
	}

	c := &cli{
		plugins,
	}

	command := c.findCommand("g")

	if command.Name() != "generate" {
		t.Errorf("Not equal")
	}
}
