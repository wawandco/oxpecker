package config

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
	return "middleware/initializer"
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

	filename := filepath.Join(f.(string), "config", "database.yml")
	err := source.Build(filename, databaseYML, n.(string))
	return err
}
