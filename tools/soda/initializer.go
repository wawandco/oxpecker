package soda

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wawandco/ox/lifecycle/new"
)

type Initializer struct{}

func (in Initializer) Name() string {
	return "soda/initializer"
}

func (in Initializer) Initialize(ctx context.Context, options new.Options) error {
	err := os.MkdirAll(filepath.Join(options.Folder, "migrations"), 0777)
	if err != nil {
		return err
	}

	readme := filepath.Join(options.Folder, "migrations", "README.md")
	content := []byte("This is the migrations folder, here live the migrations to keep the database up to date.")
	err = ioutil.WriteFile(readme, content, 0777)
	if err != nil {
		return err
	}

	return nil
}
