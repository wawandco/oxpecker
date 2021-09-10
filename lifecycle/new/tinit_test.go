package new_test

import (
	"context"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/lifecycle/new"
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

func (t *Tinit) Initialize(ctx context.Context, options new.Options) error {
	t.called = true
	t.folder = options.Folder
	t.root = options.Root
	t.name = options.Name

	return nil
}

func (t *Tinit) AfterInitialize(ctx context.Context, options new.Options) error {
	t.afterCalled = true
	t.root = options.Root

	return nil
}

func (t *Tinit) ParseFlags([]string) {}
func (t *Tinit) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("tinit", pflag.ContinueOnError)
}
