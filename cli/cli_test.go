package cli

import (
	"testing"

	"github.com/wawandco/oxpecker/lifecycle/build"
	"github.com/wawandco/oxpecker/lifecycle/dev"
	"github.com/wawandco/oxpecker/lifecycle/generate"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxpecker/tools/cli/help"
	"github.com/wawandco/oxpecker/tools/cli/version"
)

func Test_CliTestingAliaser(t *testing.T) {
	plugins := []plugins.Plugin{
		&generate.Command{},
		&build.Command{},
		&dev.Command{},
		&version.Command{},
		&help.Command{},
	}

	c := &cli{
		plugins,
	}

	tcases := []struct {
		commandAlias string
		nameExpected string
	}{
		{"g", "generate"},
		{"b", "build"},
		{"d", "dev"},
		{"v", "version"},
		{"h", "help"},
	}

	for _, ca := range tcases {
		command := c.findCommand(ca.commandAlias)
		if command.Name() != ca.nameExpected {
			t.Errorf("Not equal")
		}
	}

}
