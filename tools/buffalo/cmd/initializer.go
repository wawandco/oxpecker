package cmd

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

	filename := filepath.Join(f.(string), "cmd", n.(string), "main.go")
	err := source.Build(filename, mainGo, m.(string))

	return err
}
