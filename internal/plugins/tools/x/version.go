package x

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Command = (*VersionCommand)(nil)
)

type VersionCommand struct {
	versioner plugins.Versioner
}

func (b VersionCommand) Name() string {
	return "version"
}

func (b *VersionCommand) Run(ctx context.Context, root string, args []string) error {
	version, err := b.versioner.Version()
	if err != nil {
		return err
	}

	fmt.Printf("x version %v\n", version)

	return nil
}

func (b *VersionCommand) Receive(pl []plugins.Plugin) {
	for _, plugin := range pl {
		vr, ok := plugin.(plugins.Versioner)
		// We're looking for a versioner that for the x tool.
		if !ok || vr.ToolName() != "x" {
			continue
		}

		b.versioner = vr
		break
	}
}
