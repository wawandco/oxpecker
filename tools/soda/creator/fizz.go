package creator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"

	smartfizz "github.com/wawandco/oxpecker/tools/soda/smart_fizz"
)

// FizzCreator model struct for fizz generation files
type FizzCreator struct{}

// Name is the name of the migration type
func (f FizzCreator) Name() string {
	return "fizz"
}

// Create will create 2 .fizz files for the migration
func (f *FizzCreator) Create(dir string, args []string) error {
	name := flect.Underscore(flect.Pluralize(strings.ToLower(args[0])))

	sf := smartfizz.New(name)

	if err := sf.Generate(args); err != nil {
		return err
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	if err := f.FizzUp(dir, fileName, sf.Fizz()); err != nil {
		return err
	}

	if err := f.FizzDown(dir, fileName, sf.UnFizz()); err != nil {
		return err
	}

	return nil
}

func (f FizzCreator) FizzUp(dir, name, content string) error {
	filename := fmt.Sprintf("%s.up.fizz", name)
	path := filepath.Join(dir, filename)

	return f.createFile(path, content)
}

func (f FizzCreator) FizzDown(dir, name, content string) error {
	filename := fmt.Sprintf("%s.down.fizz", name)
	path := filepath.Join(dir, filename)

	return f.createFile(path, content)
}

func (f FizzCreator) createFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return errors.Wrap(err, "error writing file")
	}

	writer.Flush()

	return nil
}
