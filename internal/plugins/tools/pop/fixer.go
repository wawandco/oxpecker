package pop

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/paganotoni/oxpecker/internal/plugins/lifecycle/fix"
	"github.com/paganotoni/oxpecker/plugins"
)

// Err..
var (
	_ plugins.Plugin = (*Fixer)(nil)
	_ fix.Fixer      = (*Fixer)(nil)

	ErrDatabaseNotExist = errors.New(" database.yml does not exist")
)

// Fixer type ...
type Fixer struct{}

func (f Fixer) Name() string {
	return "fix database"
}

// Fix moves the file "database.yml" to
// "/config/database.yml". If the file
// already exists it ignores the oparation
func (f Fixer) Fix(ctx context.Context, root string, args []string) error {
	//search for file
	_, err := f.fileExists(".")
	if err != nil {
		return err
	}
	err = os.MkdirAll("config/", 0755)
	if err != nil {
		return err
	}
	_, err = f.fileExists("config/")
	if err == ErrDatabaseNotExist {
		err = f.moveFile()
		if err != nil {
			return err
		}
	}

	return nil
}

// moveFile moves the database.yml file to
// a config/ directory

func (f Fixer) moveFile() error {

	err := os.Rename("database.yml", "config/database.yml")
	if err != nil {
		return err
	}

	return nil
}

// fileExists search in the s directory for
// the "database.yml" file
func (f Fixer) fileExists(s string) (bool, error) {
	files, err := ioutil.ReadDir(s)
	if err != nil {
		return false, err
	}

	for _, f := range files {
		if f.Name() == "database.yml" {
			fmt.Println(f.Name() + " found")
			return true, nil
		}
	}

	return false, ErrDatabaseNotExist
}
