package help

import (
	"strings"
	"testing"

	"github.com/wawandco/oxpecker/plugins/cli/version"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxpecker/internal/plugins/tools/pop"
	"github.com/wawandco/oxpecker/internal/plugins/tools/pop/migrate"
)

func TestFindCommand(t *testing.T) {
	hp := Help{
		commands: []plugins.Command{
			&version.Version{},
		},
	}

	migrate := &migrate.Plugin{}
	pop := &pop.Plugin{}
	pop.Receive([]plugins.Plugin{
		migrate,
	})

	hp.commands = append(hp.commands, pop)

	t.Run("not enough arguments", func(*testing.T) {
		result, names := hp.findCommand([]string{"help"})
		if result != nil || names != nil {
			t.Fatal("Should be nil")
		}
	})

	t.Run("top level command", func(*testing.T) {
		result, names := hp.findCommand([]string{"help", "version"})
		expected := []string{
			"version",
		}
		if result.Name() != "version" || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

	t.Run("subcommand lookup", func(*testing.T) {
		result, names := hp.findCommand([]string{"help", "pop", "migrate"})
		expected := []string{
			"pop",
			"migrate",
		}
		ht, ok := result.(plugins.HelpTexter)
		if result.Name() != "migrate" || !ok || ht.HelpText() != migrate.HelpText() || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

	t.Run("extra args on non-subcommander", func(*testing.T) {
		result, names := hp.findCommand([]string{"help", "pop", "migrate", "other", "thing"})
		expected := []string{
			"pop",
			"migrate",
		}
		ht, ok := result.(plugins.HelpTexter)
		if result.Name() != "migrate" || !ok || ht.HelpText() != migrate.HelpText() || strings.Join(names, " ") != strings.Join(expected, " ") {
			t.Fatal("didn't find our guy")
		}
	})

}
