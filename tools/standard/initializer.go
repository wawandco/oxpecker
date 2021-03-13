package standard

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/spf13/pflag"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "standard/initializer"
}

// Initialize the go module
func (i *Initializer) Initialize(ctx context.Context, dx *sync.Map) error {
	m, ok := dx.Load("module")
	if !ok {
		return errors.New("incomplete")
	}

	f, ok := dx.Load("folder")
	if !ok {
		return errors.New("incomplete")
	}

	folder := f.(string)

	tmpl, err := template.New("go.mod").Parse(gomod)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, m)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "go.mod"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags(flags []string) {}
func (i *Initializer) Flags(flags []string) *pflag.FlagSet {
	return pflag.NewFlagSet("std/init", pflag.ContinueOnError)
}
