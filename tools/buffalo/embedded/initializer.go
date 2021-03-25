package embedded

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
	return "embedded/initializer"
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

	content, err := templates.ReadFile("templates/embeded.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(f.(string), "embed.go"), string(content), n.(string))
	return err
}
