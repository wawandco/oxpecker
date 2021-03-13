package embedded

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

	err := source.Build(filepath.Join(f.(string), "embed.go"), embedGo, n.(string))
	return err
}
