package ox

import (
	"context"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/info"
	"github.com/wawandco/oxpecker/internal/log"
	"github.com/wawandco/oxpecker/internal/source"
)

var (
	//go:embed templates/main.go.tmpl
	mainTemplate string
)

type Generator struct{}

func (g Generator) Name() string {
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
