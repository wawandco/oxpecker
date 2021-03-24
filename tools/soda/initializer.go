package soda

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Initializer struct{}

func (in Initializer) Name() string {
	return "soda/initializer"
}

func (in Initializer) Initialize(ctx context.Context) error {
	root, ok := ctx.Value("folder").(string)
	if !ok {
		return errors.New("folder needed")
	}

	err := os.MkdirAll(filepath.Join(root, "migrations"), 0777)
	if err != nil {
		return err
	}

	readme := filepath.Join(root, "migrations", "README.md")
	content := []byte("This is the migrations folder, here live the migrations to keep the database up to date.")
	err = ioutil.WriteFile(readme, content, 0777)
	if err != nil {
		return err
	}

	return nil
}
