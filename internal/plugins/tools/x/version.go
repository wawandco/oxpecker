package x

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Command = (*VersionCommand)(nil)
)

// Version command will print X version.
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

// Receive the plugins and find a Versioner for the X tool
// store it for later call to its Version function.
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
