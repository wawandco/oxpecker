package action

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

var (

	//go:embed templates
	templates embed.FS

	files = map[string]string{
		"actions_test.go.tmpl": filepath.Join("app", "actions", "actions_test.go"),
		"actions.go.tmpl":      filepath.Join("app", "actions", "actions.go"),
		"home.go.tmpl":         filepath.Join("app", "actions", "home", "home.go"),
	}

	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	entries, err := templates.ReadDir("templates")
	if err != nil {
		return err
	}

	for _, e := range entries {

		if e.IsDir() {
			continue
		}

		bt, err := fs.ReadFile(templates, filepath.Join("templates", e.Name()))
		if err != nil {
			return err
		}

		template := string(bt)
		result := files[e.Name()]
		if result == "" {
			continue
		}

		err = source.Build(filepath.Join(options.Folder, result), template, options.Module)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
