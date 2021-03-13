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

	basefolder := filepath.Join(f.(string), "app", "actions")
	folders := []string{
		filepath.Join(f.(string), "app", "actions"),
		filepath.Join(f.(string), "app", "actions", "home"),
	}

	for _, fol := range folders {
		err := os.MkdirAll(fol, 0777)
		if err != nil {
			return err
		}
	}

	instructions := []struct {
		file     string
		template string
	}{
		{file: filepath.Join(basefolder, "actions.go"), template: actionsGo},
		{file: filepath.Join(basefolder, "actions_test.go"), template: actionsTestGo},
		{file: filepath.Join(basefolder, "home", "home.go"), template: homeGo},
	}

	data := struct {
		Module string
	}{
		Module: m.(string),
	}

	for _, ins := range instructions {
		tmpl, err := template.New(ins.file).Parse(ins.template)
		if err != nil {
			return err
		}

		sbf := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(sbf, data)

		if err != nil {
			return err
		}

		err = ioutil.WriteFile(ins.file, sbf.Bytes(), 0777)
		if err != nil {
			return err
		}
	}

	fmt.Printf("[info] Created app/actions/actions.go\n")
	fmt.Printf("[info] Created app/actions/actions_test.go\n")
	fmt.Printf("[info] Created app/actions/home/home.go\n")

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
