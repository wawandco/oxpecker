package creator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

// SQLCreator model struct for fizz generation files
type SQLCreator struct{}

// Name is the name of the migration type
func (s SQLCreator) Name() string {
	return "sql"
}

// Create will create 2 .sql empty files for the migration
func (s *SQLCreator) Create(dir string, args []string) error {
	name := flect.Underscore(flect.Pluralize(strings.ToLower(args[0])))
	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	if err := s.createFile(dir, fileName, "up"); err != nil {
		return err
	}

	if err := s.createFile(dir, fileName, "down"); err != nil {
		return err
	}

	return nil
}

func (s *SQLCreator) createFile(dir, name, runFlag string) error {
	fileName := fmt.Sprintf("%s.%s.sql", name, runFlag)
	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	return err
}
