package action

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

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

func (i *Initializer) Initialize(ctx context.Context, dx *sync.Map) error {
	m, ok := dx.Load("module")
	if !ok {
		return ErrIncompleteArgs
	}

	f, ok := dx.Load("folder")
	if !ok {
		return ErrIncompleteArgs
	}

	folder := filepath.Join(f.(string), "app", "actions")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New("actions_test.go").Parse(actionsTestGo)
	if err != nil {
		return err
	}

	data := struct {
		Module string
	}{
		Module: m.(string),
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, data)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "actions_test.go"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	tmpl, err = template.New("actions.go").Parse(actionsGo)
	if err != nil {
		return err
	}

	data = struct {
		Module string
	}{
		Module: m.(string),
	}

	sbf = bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, data)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(folder, "actions.go"), sbf.Bytes(), 0777)
	if err != nil {
		return err
	}

	fmt.Printf("[info] Created: \n-- app/actions/actions.go\n-- app/actions/actions_test.go\n")
	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
