package docker

import (
	"context"
	"embed"
	"errors"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

var (
	//go:embed templates
	templates embed.FS
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "docker/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {

	folder, ok := ctx.Value("folder").(string)
	if !ok {
		return errors.New("folder needed")
	}

	files := []struct {
		path     string
		template string
	}{
		{filepath.Join(folder, ".dockerignore"), "templates/dot-dockerignore.tmpl"},
		{filepath.Join(folder, "Dockerfile"), "templates/Dockerfile.tmpl"},
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
