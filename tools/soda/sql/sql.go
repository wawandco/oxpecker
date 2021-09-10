package sql

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/log"
)

// Creator model struct for fizz generation files
type Creator struct{}

// Name is the name of the migration type
func (s Creator) Name() string {
	return "sql"
}

// Creates a type or not
func (f Creator) Creates(mtype string) bool {
	return mtype == "sql"
}

// Create will create 2 .sql empty files for the migration
func (s Creator) Create(dir, name string, args []string) error {
	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	if err := s.createFile(dir, fileName, "up"); err != nil {
		return err
	}

	if err := s.createFile(dir, fileName, "down"); err != nil {
		return err
	}

	log.Infof("generated: %s/%s.up.sql", dir, fileName)
	log.Infof("generated: %s/%s.down.sql", dir, fileName)

	return nil
}

func (s Creator) createFile(dir, name, runFlag string) error {
	fileName := fmt.Sprintf("%s.%s.sql", name, runFlag)
	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}

	defer file.Close()

	return err
}
