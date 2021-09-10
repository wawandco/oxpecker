package action

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/source"
)

type Generator struct {
	name     string
	filename string
	dir      string
}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "buffalo/generate-action"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
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
	path := filepath.Join(g.dir, g.filename+".go")
	data := struct {
		Name string
	}{
		Name: g.name,
	}
	actionTemplate, err := g.callTemplate("action.go.tmpl")
	if err != nil {
		return errors.Wrap(err, "error calling template")
	}

	err = source.Build(path, actionTemplate, data)
	if err != nil {
		return errors.Wrap(err, "error generating action")
	}

	return nil
}

func (g Generator) createActionTestFile() error {
	path := filepath.Join(g.dir, g.filename+"_test.go")
	data := struct {
		Name string
	}{
		Name: g.name,
	}

	actionTestTemplate, err := g.callTemplate("action_test.go.tmpl")
	if err != nil {
		return errors.Wrap(err, "error calling template")
	}

	err = source.Build(path, actionTestTemplate, data)
	if err != nil {
		return errors.Wrap(err, "error generating action")
	}

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func (g Generator) callTemplate(name string) (string, error) {
	bt, err := fs.ReadFile(templates, filepath.Join("templates", name))
	if err != nil {
		return "", nil
	}

	return string(bt), nil
}
