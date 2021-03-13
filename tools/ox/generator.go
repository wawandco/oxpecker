package ox

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wawandco/oxpecker/internal/info"
	"github.com/wawandco/oxpecker/internal/source"
)

type Generator struct{}

func (g Generator) Name() string {
	return "ox"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	file := filepath.Join("cmd", "ox", "main.go")
	if _, err := os.Stat(file); err == nil {
		fmt.Println("[info] skipping file generation because ox/main.go exists.")
		return nil
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	module, err := info.ModuleName()
	if err != nil {
		return err
	}

	data := struct {
		Name   string
		Module string
	}{
		Name:   name,
		Module: module,
	}

	err = source.Build(file, mainTemplate, data)
	return err
}
