package help

import (
	"context"
	"strings"
	"testing"

	"github.com/wawandco/ox/plugins"
)

func TestFindCommand(t *testing.T) {
	hp := Command{
		commands: []plugins.Command{},
	}

	migrate := &subPl{}
	pop := &testPlugin{}
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
		result, names := hp.findCommand([]string{"help", "pop"})
		expected := []string{
			"pop",
		}
		if result.Name() != "pop" || strings.Join(names, " ") != strings.Join(expected, " ") {
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

type testPlugin struct {
	subcommands []plugins.Command
}

func (tp testPlugin) Name() string {
	return "pop"
}

func (tp testPlugin) ParentName() string {
	return ""
}

func (tp testPlugin) HelpText() string {
	return "pop help text"
}

func (tp *testPlugin) Run(ctx context.Context, root string, args []string) error {
	return nil
}

func (tp *testPlugin) Receive(pls []plugins.Plugin) {
	for _, pl := range pls {
		c, ok := pl.(plugins.Command)
		if !ok || c.ParentName() != tp.Name() {
			continue
		}

		tp.subcommands = append(tp.subcommands, c)
	}
}

func (tp *testPlugin) Subcommands() []plugins.Command {
	return tp.subcommands
}

type subPl struct{}

func (tp subPl) Name() string {
	return "migrate"
}

func (tp subPl) ParentName() string {
	return "pop"
}

func (tp subPl) HelpText() string {
	return "migrate help text"
}

func (tp subPl) Run(ctx context.Context, root string, args []string) error {
	return nil
}
