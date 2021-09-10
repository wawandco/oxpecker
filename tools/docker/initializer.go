package docker

import (
	"context"
	"embed"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

var (
	//go:embed templates
	templates embed.FS
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "docker/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	files := []struct {
		path     string
		template string
	}{
		{filepath.Join(options.Folder, ".dockerignore"), "templates/dot-dockerignore.tmpl"},
		{filepath.Join(options.Folder, "Dockerfile"), "templates/Dockerfile.tmpl"},
	}

	for _, f := range files {
		content, err := templates.ReadFile(f.template)
		if err != nil {
			return err
		}

		err = source.Build(f.path, string(content), nil)
		if err != nil {
			return err
		}
	}

	return nil
}
