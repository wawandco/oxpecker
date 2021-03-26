package template

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
	return "template/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {
	n := ctx.Value("name")
	if n == nil {
		return ErrIncompleteArgs
	}

	f := ctx.Value("folder")
	if f == nil {
		return ErrIncompleteArgs
	}

	layout, err := templates.ReadFile("templates/layout.html.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(f.(string), "app", "templates", "application.plush.html")
	err = source.Build(filename, string(layout), n.(string))
	if err != nil {
		return err
	}

	home, err := templates.ReadFile("templates/home.html.tmpl")
	if err != nil {
		return err
	}

	filename = filepath.Join(f.(string), "app", "templates", "home", "index.plush.html")
	err = source.Build(filename, string(home), n.(string))

	return err
}
