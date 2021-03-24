package assets

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/source"
)

var (
	//go:embed templates
	templates embed.FS

	// files that will be created on the new app as
	// part of the asset pipeline
	files = map[string]string{
		"application.js.tmpl":    filepath.Join("app", "assets", "js", "application.js"),
		"application.scss.tmpl":  filepath.Join("app", "assets", "css", "application.scss"),
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

func (i Initializer) Initialize(ctx context.Context) error {
	module := ctx.Value("module")
	if module == nil {
		return errors.New("module is needed")
	}

	folder := ctx.Value("folder")
	if folder == nil {
		return errors.New("folder is needed")
	}

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

		err = source.Build(filepath.Join(folder.(string), result), template, module)
		if err != nil {
			return err
		}
	}

	return nil
}
