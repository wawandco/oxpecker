package fizz

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"time"

	"github.com/wawandco/ox/internal/source"
)

var (
	//go:embed templates/content.fizz.tmpl
	fizzTemplate string
)

// FizzCreator model struct for fizz generation files
type Creator struct{}

// Name is the name of the migration type
func (f Creator) Name() string {
	return "fizz"
}

// Creates a type or not
func (f Creator) Creates(mtype string) bool {
	return mtype == "fizz"
}

// Create will create 2 .fizz files for the migration
func (f Creator) Create(dir, name string, args []string) error {
	g := generators.GeneratorFor(name)

	up, down, err := g.GenerateFizz(name, args)
	if err != nil {
		return err
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	upPath := filepath.Join(dir, fileName+".up.fizz")
	downPath := filepath.Join(dir, fileName+".down.fizz")

	// Build Up Fizz
	if err := source.Build(upPath, fizzTemplate, up); err != nil {
		return err
	}

	// Build Down Fizz
	if err := source.Build(downPath, fizzTemplate, down); err != nil {
		return err
	}

	return nil
}
