package ox

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/wawandco/oxpecker/internal/info"
)

type Generator struct{}

func (g Generator) Name() string {
	return "ox"
}

func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	file := filepath.Join("cmd", "ox", "main.go")
	if _, err := os.Stat(file); err == nil {
		fmt.Println("file exists!")
		return nil
	}

	if _, err := os.Stat(file); err != nil {
		//Folder does not exist, we proceed to create it
		err = os.MkdirAll(filepath.Dir(file), 0644)
		if err != nil {
			return err
		}
	}

	tmpl, err := template.New("main.go").Parse(mainTemplate)
	if err != nil {
		return err
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	data := struct {
		Name string
	}{
		Name: name,
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return err
	}

	return ioutil.WriteFile(file, tpl.Bytes(), 0644)
}
