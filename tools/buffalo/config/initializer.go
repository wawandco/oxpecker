package config

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

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "middleware/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	template, err := templates.ReadFile("templates/postgres.database.yml.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "config", "database.yml")
	err = source.Build(filename, string(template), options.Name)
	return err
}
