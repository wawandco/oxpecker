package assets

import (
	"context"
	"embed"
	"io/fs"
	"path/filepath"

	"github.com/wawandco/ox/internal/source"
	"github.com/wawandco/ox/lifecycle/new"
)

var (
	//go:embed templates
	templates embed.FS

	// files that will be created on the new app as
	// part of the asset pipeline
	files = map[string]string{
		"application.js.tmpl":    filepath.Join("app", "assets", "js", "application.js"),
		"application.scss.tmpl":  filepath.Join("app", "assets", "css", "application.scss"),
		"buffalo.svg.tmpl":       filepath.Join("app", "assets", "images", "buffalo.svg"),
		"dot-babelrc.tmpl":       ".babelrc",
		"package.json.tmpl":      "package.json",
		"webpack.config.js.tmpl": "webpack.config.js",
		"postcss.config.js.tmpl": "postcss.config.js",
		"robots.txt.tmpl":        filepath.Join("public", "robots.txt"),
	}
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "assets/initializer"
}

func (i Initializer) Initialize(ctx context.Context, options new.Options) error {
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
