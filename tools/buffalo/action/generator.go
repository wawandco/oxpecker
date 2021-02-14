package action

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

type Generator struct {
	name     string
	filename string
	dir      string
}

func (g Generator) Name() string {
	return "action"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate action [name]`")
	}

	dirPath := filepath.Join(root, "app", "actions")
	if !g.exists(dirPath) {
		err := os.MkdirAll(filepath.Dir(dirPath), 0755)
		if err != nil {
			return (err)
		}
	}

	g.name = flect.Singularize(args[2])
	g.filename = flect.Singularize(flect.Underscore(args[2]))
	g.dir = dirPath

	if g.exists(filepath.Join(g.dir, g.filename+".go")) {
		return errors.Errorf("action file already exists")
	}

	if err := g.generateActionFiles(args[3:]); err != nil {
		return err
	}

	fmt.Printf("[info] Action generated in: \n-- app/actions/%s.go\n-- app/actions/%s_test.go\n", g.name, g.name)

	return nil
}

func (g Generator) generateActionFiles(args []string) error {
	if err := g.createActionFile(args); err != nil {
		return errors.Wrap(err, "creating action file")
	}

	if err := g.createActionTestFile(); err != nil {
		return errors.Wrap(err, "creating action test file")
	}

	return nil
}

func (g Generator) createActionFile(args []string) error {
	filename := g.filename + ".go"
	path := filepath.Join(g.dir, filename)
	data := struct {
		Name string
	}{
		Name: g.name,
	}

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(actionTemplate)
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

func (g Generator) createActionTestFile() error {
	filename := g.filename + "_test.go"
	path := filepath.Join(g.dir, filename)
	data := struct {
		Name string
	}{
		Name: g.name,
	}

	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(actionTestTemplate)
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

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
