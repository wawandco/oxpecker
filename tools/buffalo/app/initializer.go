package app

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

	n := ctx.Value("name")
	if n == nil {
		return ErrIncompleteArgs
	}

	data := struct {
		Module, Name string
	}{
		Module: m.(string),
		Name:   n.(string),
	}

	appGo, err := templates.ReadFile("templates/app.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(f.(string), "app", "app.go"), string(appGo), data)
	if err != nil {
		return err
	}

	routesGo, err := templates.ReadFile("templates/routes.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(f.(string), "app", "routes.go"), string(routesGo), data)
	if err != nil {
		return err
	}

	return nil
}
