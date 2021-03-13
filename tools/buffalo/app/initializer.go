package app

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/pflag"
)

var (
	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {
	m := ctx.Value("module")
	if m == nil {
		return ErrIncompleteArgs
	}

	f := ctx.Value("folder")
	if f == nil {
		return ErrIncompleteArgs
	}

	n := ctx.Value("name")
	if n == nil {
		return ErrIncompleteArgs
	}

	folder := filepath.Join(f.(string), "app")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("app.go").Parse(appGo)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, struct {
		Module, Name string
	}{
		Module: m.(string),
		Name:   n.(string),
	})
	if err != nil {
		return err
	}

	path := filepath.Join(folder, "app.go")
	err = ioutil.WriteFile(path, sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	tmpl, err = template.New("routes.go").Parse(routesGo)
	if err != nil {
		return err
	}

	sbf = bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, struct {
		Module, Name string
	}{
		Module: m.(string),
		Name:   n.(string),
	})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "routes.go"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
