package ox

import (
	"context"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/internal/source"
)

var (
	//go:embed templates/main.go.tmpl
	mainTemplate string
)

type Generator struct{}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "ox/generate-cli-main"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "ox"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	file := filepath.Join("cmd", "ox", "main.go")
	if _, err := os.Stat(file); err == nil {
		log.Info("skipping file generation because ox/main.go exists.")
		return nil
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	data := struct {
		Name   string
		Module string
	}{
		Name:   name,
		Module: info.ModuleName(),
	}

	err = source.Build(file, mainTemplate, data)
	return err
}
