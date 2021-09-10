package template

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
	return "template/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	layout, err := templates.ReadFile("templates/layout.html.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "app", "templates", "application.plush.html")
	err = source.Build(filename, string(layout), options.Name)
	if err != nil {
		return err
	}

	home, err := templates.ReadFile("templates/home.html.tmpl")
	if err != nil {
		return err
	}

	filename = filepath.Join(options.Folder, "app", "templates", "home", "index.plush.html")
	err = source.Build(filename, string(home), options.Name)

	return err
}
