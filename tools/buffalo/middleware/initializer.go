package middleware

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
	return "middleware/initializer"
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

	folder := filepath.Join(f.(string), "app", "middleware")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("middleware.go").Parse(middlewareGo)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, struct {
		Module string
	}{
		Module: m.(string),
	})
	if err != nil {
		return err
	}

	path := filepath.Join(folder, "middleware.go")
	err = ioutil.WriteFile(path, sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
