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
	called      bool

	root   string
	folder string
	name   string
}

func (t Tinit) Name() string { return "tinit" }

func (t *Tinit) Initialize(ctx context.Context) error {
	t.called = true

	if f := ctx.Value("root"); f != nil {
		t.folder = f.(string)
	}

	if f := ctx.Value("folder"); f != nil {
		t.folder = f.(string)
	}

	if f := ctx.Value("name"); f != nil {
		t.name = f.(string)
	}

	return nil
}

func (t *Tinit) AfterInitialize(ctx context.Context) error {
	t.afterCalled = true

	if f := ctx.Value("root"); f != nil {
		t.root = f.(string)
	}

	return nil
}

func (t *Tinit) ParseFlags([]string) {}
func (t *Tinit) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("tinit", pflag.ContinueOnError)
}
