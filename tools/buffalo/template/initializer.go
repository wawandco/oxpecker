package template

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

var (
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

	filename := filepath.Join(f.(string), "app", "templates", "application.plush.html")
	err := source.Build(filename, layout, n.(string))
	if err != nil {
		return err
	}

	filename = filepath.Join(f.(string), "app", "templates", "home", "index.plush.html")
	err = source.Build(filename, home, n.(string))

	return err
}
