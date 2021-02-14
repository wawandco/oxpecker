package new_test

import (
	"context"
	"sync"

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

func (t *Tinit) Initialize(ctx context.Context, dx *sync.Map) error {
	t.called = true

	if f, ok := dx.Load("root"); ok {
		t.folder = f.(string)
	}

	if f, ok := dx.Load("folder"); ok {
		t.folder = f.(string)
	}

	if f, ok := dx.Load("name"); ok {
		t.name = f.(string)
	}

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
