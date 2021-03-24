package grift

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "grift/initializer"
}

func (i Initializer) Initialize(ctx context.Context) error {
	module := ctx.Value("module")
	if module == nil {
		return errors.New("module name needed")
	}

	folder := ctx.Value("folder")
	if module == nil {
		return errors.New("folder name needed")
	}

	err := source.Build(filepath.Join(folder.(string), "app", "tasks", "tasks.go"), tasksGo, module)
	return err
}
