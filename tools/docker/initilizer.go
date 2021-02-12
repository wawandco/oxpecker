package docker

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/wawandco/oxplugins/internal/info"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "docker/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, arg []string) error {

	rootDoc := root + "/.dockerignore"
	rootFile := filepath.Join(root, "Dockerfile")
	_, err := os.Create(rootFile)
	if err != nil {
		return err
	}

	contentIgnore := `.git
	node_modules/
	*.log
	vendor/
	public/assets
	tmp/
	bin`

	_, err = os.Stat(rootDoc)
	if err == nil {

		fmt.Println("dockerignore file already exist ")
		return nil

	}
	if os.IsNotExist(err) {

		// create file if it does not exist
		file, err := os.Create(rootDoc)

		if err != nil {
			return (err)
		}

		_, err = os.OpenFile(rootDoc, os.O_RDWR, 0644)
		if err != nil {
			return (err)
		}

		_, err = file.WriteString(contentIgnore)
		if err != nil {
			return (err)
		}

		file.Close()

		return nil

	}

	_, err = os.Stat(rootFile)
	if err == nil {

		fmt.Println("Dockerfile file already exist ")
		return nil
	}
	if os.IsNotExist(err) {
		_, err := os.Create(rootFile)
		if err != nil {
			return (err)
		}

		tmpl, err := template.New("Dockerfile").Parse(dockerTemplate)

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

		err = ioutil.WriteFile(rootFile, tpl.Bytes(), 0655)

		if err != nil {
			return err
		}

	}
	return err
}
