package app

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
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	appGo, err := templates.ReadFile("templates/app.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(options.Folder, "app", "app.go"), string(appGo), options)
	if err != nil {
		return err
	}

	routesGo, err := templates.ReadFile("templates/routes.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(options.Folder, "app", "routes.go"), string(routesGo), options)
	if err != nil {
		return err
	}

	return nil
}
