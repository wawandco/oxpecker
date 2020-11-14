package x

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

var (
	ErrModuleNameNeeded   = errors.New("module name needed")
	ErrModuleNameNotFound = errors.New("module name not found")
	ErrFileMainNotExist   = errors.New("main.go file does not exist")
)

// Fixer is in charge of performing a Fix operation
// that moves the main.go to cmd/[name-of-the-module]/main.go
type Fixer struct{}

// Fix does the main.go magic
// - Determine if the file exists
// - Determine if there is a go.mod
// - Determine the name of the module (last part when slicing go.mod by /)
// - Create folder
// - Copy/move main.go to that folder
func (f Fixer) Fix() error {
	_, err := f.fileExists()
	if err != nil {
		return err
	}

	base, err := f.findModuleName()
	if err != nil {
		return err
	}

	err = f.moveFile(base)
	if err != nil {
		return err
	}

	return nil
}

func (f Fixer) moveFile(s string) error {
	if s == "" {
		return ErrModuleNameNeeded
	}

	name := "main.go"
	s = "cmd/" + s
	err := os.MkdirAll(s, 0755)
	if err != nil {
		return err
	}

	finLoc := s + "/" + name
	err = os.Rename(name, finLoc)
	if err != nil {
		return err
	}

	return nil
}

// Look for go.mod and extract the last part.
func (f Fixer) findModuleName() (string, error) {
	mp := "go.mod"
	file, err := ioutil.ReadFile(mp)
	if err != nil {
		return "", err
	}

	base := filepath.Base(modfile.ModulePath(file))

	if base != "." {
		return base, nil
	}
	if base == "." {
		return "", ErrModuleNameNotFound
	}
	return "", ErrModuleNameNotFound
}

func (f Fixer) fileExists() (bool, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return false, err
	}

	for _, f := range files {
		if f.Name() == "main.go" {
			fmt.Println(f.Name() + " found")
			return true, nil
		}
	}

	return false, ErrFileMainNotExist
}
