package grift

import (
	"context"
	"embed"
	"errors"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

var (
	//go:embed templates
	templates embed.FS
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "grift/initializer"
}

func (i Initializer) Initialize(ctx context.Context) error {
	module := ctx.Value("module")
	if module == nil {
		return errors.New("module name needed")
	}

	folder := ctx.Value("folder")
	if module == nil {
		return errors.New("folder name needed")
	}

	content, err := templates.ReadFile("templates/grift.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(folder.(string), "app", "tasks", "tasks.go"), string(content), module)
	return err
}
