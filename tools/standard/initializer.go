package standard

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// Initialize the go module
func (i *Initializer) Initialize(ctx context.Context) error {
	m := ctx.Value("module")
	if m == nil {
		return errors.New("incomplete")
	}

	f := ctx.Value("folder")
	if f == nil {
		return errors.New("incomplete")
	}

	err := source.Build(filepath.Join(f.(string), "go.mod"), gomod, m)
	return err
}
