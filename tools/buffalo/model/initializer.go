package model

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	tmpl, err := templates.ReadFile("templates/models.go.tmpl")
	if err != nil {
		return err
	}

	filename := filepath.Join(options.Folder, "app", "models", "models.go")
	err = source.Build(filename, string(tmpl), options.Module)
	if err != nil {
		return err
	}

	tmpl, err = templates.ReadFile("templates/models_test.go.tmpl")
	if err != nil {
		return err
	}

	filename = filepath.Join(options.Folder, "app", "models", "models_test.go")
	err = ioutil.WriteFile(filename, tmpl, 0777)
	if err != nil {
		return err
	}

	return nil
}
