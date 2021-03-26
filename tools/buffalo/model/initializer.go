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

	tmpl, err := templates.ReadFile("templates/models.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(f.(string), "app", "models", "models.go")
	err = source.Build(filename, string(tmpl), m.(string))
	if err != nil {
		return err
	}

	tmpl, err = templates.ReadFile("templates/models_test.go.tmpl")
	if err != nil {
		return err
	}

	filename = filepath.Join(f.(string), "app", "models", "models_test.go")
	err = ioutil.WriteFile(filename, tmpl, 0777)
	if err != nil {
		return err
	}

	return nil
}
