package model

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/flect/name"
	"github.com/pkg/errors"
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

	tmpl, err := template.New(filename).Parse(modelTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}

func (m Model) createModelTestFile() error {
	filename := m.filename + "_test.go"
	path := filepath.Join(m.dir, filename)
	data := opts{
		Original: m.name,
		Name:     name.New(m.name),
	}

	tmpl, err := template.New(filename).Parse(modelTestTemplate)
	if err != nil {
		return errors.Wrap(err, "parsing new template error")
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrap(err, "executing new template error")
	}

	err = ioutil.WriteFile(path, tpl.Bytes(), 0655)
	if err != nil {
		return errors.Wrap(err, "writing new template error")
	}

	return nil
}
