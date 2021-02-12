package standard

import (
	"context"

	"github.com/spf13/pflag"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// - Initializes module based on args[0]
// - Creates cmd/name/main.go
func (i *Initializer) AfterInitialize(ctx context.Context, root string, args []string) error {
	return nil
}

func (i *Initializer) ParseFlags(flags []string) {}
func (i *Initializer) Flags(flags []string) *pflag.FlagSet {
	return pflag.NewFlagSet("std/init", pflag.ContinueOnError)
}
