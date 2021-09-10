package render

import (
	"context"
	"embed"
	"errors"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

var (
	//go:embed templates
	templates embed.FS

	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	renderGo, err := templates.ReadFile("templates/render.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "app", "render", "render.go")
	err = source.Build(filename, string(renderGo), options.Module)

	return err
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
