package action

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/internal/source"
)

var (
	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {
	m := ctx.Value("module")
	if m == nil {
		return ErrIncompleteArgs
	}

	f := ctx.Value("folder")
	if f == nil {
		return ErrIncompleteArgs
	}

	basefolder := filepath.Join(f.(string), "app", "actions")
	instructions := []struct {
		file     string
		template string
	}{
		{file: filepath.Join(basefolder, "actions.go"), template: actionsGo},
		{file: filepath.Join(basefolder, "actions_test.go"), template: actionsTestGo},
		{file: filepath.Join(basefolder, "home", "home.go"), template: homeGo},
	}

	for _, ins := range instructions {
		err := source.Build(ins.file, ins.template, m.(string))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
