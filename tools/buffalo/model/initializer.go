package model

import (
	"context"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

var (
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

	filename := filepath.Join(f.(string), "app", "models", "models.go")
	err := source.Build(filename, modelsBaseTemplate, m.(string))
	if err != nil {
		return err
	}

	filename = filepath.Join(f.(string), "app", "models", "models_test.go")
	err = ioutil.WriteFile(filename, []byte(modelsTestBaseTemplate), 0777)
	if err != nil {
		return err
	}

	return nil
}
