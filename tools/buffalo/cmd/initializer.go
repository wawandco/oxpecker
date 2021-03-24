package cmd

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
	return "cmd/initializer"
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

	content, err := templates.ReadFile("templates/main.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(f.(string), "cmd", n.(string), "main.go")
	err = source.Build(filename, string(content), m.(string))

	return err
}
