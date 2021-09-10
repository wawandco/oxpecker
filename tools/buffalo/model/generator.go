package model

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/log"
)

var (
	ErrModelAlreadyExists error = errors.New("model already exists")
	ErrModelDirNotFound   error = errors.New("models folder do not exists on your buffalo app, please ensure the folder exists in order to proceed")
)

// Generator allows to identify model as a plugin
type Generator struct{}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "buffalo/generate-model"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "model"
}

// Generate generates an empty [name].plush.html file
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate model [name]`")
	}

	dirPath := filepath.Join(root, "app", "models")
	if !g.exists(dirPath) {
		return ErrModelDirNotFound
	}

	filename := flect.Singularize(flect.Underscore(args[2]))

	if g.exists(filepath.Join(dirPath, filename+".go")) {
		return ErrModelAlreadyExists
	}

	model := New(dirPath, args[2], args[3:])
	if err := model.Create(); err != nil {
		return errors.Wrap(err, "creating models error")
	}

	log.Infof("Model generated in: \n-- app/models/%s.go\n-- app/models/%s_test.go\n", filename, filename)

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
