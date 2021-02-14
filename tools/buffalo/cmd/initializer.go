package cmd

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/spf13/pflag"
)

var (
	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "cmd/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, dx *sync.Map) error {
	m, ok := dx.Load("module")
	if !ok {
		return ErrIncompleteArgs
	}

	f, ok := dx.Load("folder")
	if !ok {
		return ErrIncompleteArgs
	}

	n, ok := dx.Load("name")
	if !ok {
		return ErrIncompleteArgs
	}

	folder := filepath.Join(f.(string), "cmd", n.(string))
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("main.go").Parse(mainGo)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, m.(string))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "main.go"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
