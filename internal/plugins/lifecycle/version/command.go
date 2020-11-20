package version

import (
	"context"
	"fmt"

	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Command = (*Command)(nil)
)

type Command struct {
	Version string
}

func (b Command) Name() string {
	return "version"
}

func (b *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Println("Running [version] command")

	fmt.Println("Current version of x: ", b.Version)

	return nil
}
