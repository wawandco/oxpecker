package model

import (
	"embed"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/flect/name"
	"github.com/wawandco/ox/internal/source"
)

var (
	//go:embed templates
	templates embed.FS
)

type Model struct {
	Attrs []attr

	dir      string
	name     string
	filename string
}

func New(dirPath, name string, args []string) Model {
	return Model{
		Attrs: buildAttrs(args),

		dir:      dirPath,
		filename: flect.Singularize(flect.Underscore(name)),
		name:     flect.Singularize(name),
	}
}

func (m Model) Create() error {
	if err := m.createModelFile(); err != nil {
		return err
	}

	if err := m.createModelTestFile(); err != nil {
		return err
	}

	return nil
}

func (m Model) createModelFile() error {
	filename := m.filename + ".go"
	path := filepath.Join(m.dir, filename)
	data := opts{
		Original: m.name,
		Name:     name.New(m.name),
		Attrs:    m.Attrs,
		Imports:  buildImports(m.Attrs),
	}

	tmpl, err := templates.ReadFile("templates/newmodel.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(path, string(tmpl), data)
	return err
}

func (m Model) createModelTestFile() error {
	filename := m.filename + "_test.go"
	path := filepath.Join(m.dir, filename)
	data := opts{
		Original: m.name,
		Name:     name.New(m.name),
	}

	tmpl, err := templates.ReadFile("templates/newmodel_test.go.tmpl")
	if err != nil {
		return err
	}

	err = source.Build(path, string(tmpl), data)
	return err
}
