package x

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Command        = (*Version)(nil)
	_ plugins.PluginReceiver = (*Version)(nil)
)

// Version command will print X version.
type Version struct {
	versioner plugins.Versioner
}

func (b Version) Name() string {
	return "version"
}

// Run prints the version of the X cli by using the
// Versioner in the command, one for the x tool.
func (b *Version) Run(ctx context.Context, root string, args []string) error {
	version, err := b.versioner.Version()
	if err != nil {
		return err
	}

	fmt.Printf("x version %v\n", version)

	return nil
}

// Receive the plugins and find a Versioner for the X tool
// store it for later call to its Version function.
func (b *Version) Receive(pl []plugins.Plugin) {
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
