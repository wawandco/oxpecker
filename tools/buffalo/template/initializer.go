package template

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
	return "template/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, dx *sync.Map) error {
	n, ok := dx.Load("name")
	if !ok {
		return ErrIncompleteArgs
	}

	f, ok := dx.Load("folder")
	if !ok {
		return ErrIncompleteArgs
	}

	folder := filepath.Join(f.(string), "app", "templates", "home")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("application.plush.html").Parse(layout)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, n.(string))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(f.(string), "app", "templates", "application.plush.html"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	tmpl, err = template.New("index.plush.html").Parse(home)
	if err != nil {
		return err
	}

	sbf = bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, n.(string))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(f.(string), "app", "templates", "home", "index.plush.html"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
