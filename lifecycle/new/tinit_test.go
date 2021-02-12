package new_test

import (
	"context"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/lifecycle/new"
)

var _ new.Initializer = (*Tinit)(nil)
var _ new.AfterInitializer = (*Tinit)(nil)

type Tinit struct {
	afterCalled bool
	root        string
	called      bool
}

func (t Tinit) Name() string { return "tinit" }

func (t *Tinit) Initialize(ctx context.Context, root string, args []string) error {
	t.called = true
	t.root = root

	return nil
}

func (t *Tinit) AfterInitialize(ctx context.Context, root string, args []string) error {
	t.afterCalled = true
	t.root = root

	return nil
}

func (t *Tinit) ParseFlags([]string) {}
func (t *Tinit) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("tinit", pflag.ContinueOnError)
}
