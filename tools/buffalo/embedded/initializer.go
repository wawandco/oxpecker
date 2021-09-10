package embedded

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

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "embedded/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	content, err := templates.ReadFile("templates/embeded.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(filepath.Join(options.Folder, "embed.go"), string(content), options.Name)
	return err
}
