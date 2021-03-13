package model

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

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

	folder := filepath.Join(f.(string), "app", "models")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("models.go").Parse(modelsBaseTemplate)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, struct{ Module string }{
		Module: m.(string),
	})

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "models.go"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "models_test.go"), []byte(modelsTestBaseTemplate), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
