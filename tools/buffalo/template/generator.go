package template

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/log"
)

// Generator allows to identify template as a plugin
type Generator struct{}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "buffalo/generate-template"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "template"
}

// Generate generates an empty [name].plush.html file
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return errors.Errorf("no name specified, please use `ox generate template [name]`")
	}

	if err := g.generateTemplate(root, args[2]); err != nil {
		return err
	}

	log.Infof("Template generated in app/templates/%s.plush.html \n", args[2])

	return nil
}

func (g Generator) generateTemplate(root, filename string) error {
	dirpath := filepath.Join(root, "app", "templates")

	if !g.exists(dirpath) {
		return errors.Errorf("folder '%s' do not exists on your buffalo app, please ensure the folder exists in order to proceed", dirpath)
	}

	tmpl := g.generateFilePath(dirpath, filename)
	if g.exists(tmpl) {
		return errors.Errorf("template already exists")
	}

	if err := os.MkdirAll(filepath.Dir(tmpl), 0755); err != nil {
		return errors.Wrap(err, "error creating subfolders")
	}

	file, err := os.Create(tmpl)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	return nil
}

// generateFilePath translates the required path to create the file properly
func (g Generator) generateFilePath(dirPath, filename string) string {
	base := strings.Split(filename, ".")[0]
	file := base + ".plush.html"

	return filepath.Join(dirPath, file)
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
