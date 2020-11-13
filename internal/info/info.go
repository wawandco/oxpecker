package info

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

var ErrModuleNameNotFound = errors.New("module name not found")

// BuildName extracts the last part of the module by splitting on `/`
// this last part is useful for name of the binary and other things.
func BuildName() (string, error) {
	content, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	path := modfile.ModulePath(content)
	name := filepath.Base(path)

	if name == "." {
		return "", ErrModuleNameNotFound
	}

	return name, nil
}
