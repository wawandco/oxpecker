package standard

import (
	"context"
	"sync"

	"github.com/spf13/pflag"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// Initialize the go module
func (i *Initializer) Initialize(ctx context.Context, data *sync.Map) error {
	return nil
}

func (i *Initializer) ParseFlags(flags []string) {}
func (i *Initializer) Flags(flags []string) *pflag.FlagSet {
	return pflag.NewFlagSet("std/init", pflag.ContinueOnError)
}
