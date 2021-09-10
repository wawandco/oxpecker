package grift

import (
	"context"
	"embed"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

var (
	//go:embed templates
	templates embed.FS
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "grift/initializer"
}

func (i Initializer) Initialize(ctx context.Context, options new.Options) error {
	content, err := templates.ReadFile("templates/grift.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(options.Folder, "app", "tasks", "tasks.go"), string(content), options.Module)
	return err
}
