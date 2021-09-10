package cmd

import (
	"context"
	"embed"
	"errors"
	"path/filepath"

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
	return "cmd/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	content, err := templates.ReadFile("templates/main.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "cmd", options.Name, "main.go")
	err = source.Build(filename, string(content), options.Module)

	return err
}
