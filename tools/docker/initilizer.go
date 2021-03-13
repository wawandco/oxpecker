package docker

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "docker/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, arg []string) error {

	files := []struct {
		path    string
		content string
	}{
		{filepath.Join(root, ".dockerignore"), ignoreTemplate},
		{filepath.Join(root, "Dockerfile"), dockerTemplate},
	}

	for _, f := range files {
		_, err := os.Stat(f.path)
		if err == nil {
			fmt.Printf("[info] `%v` already exist, skipping\n", f.path)

			continue
		}

		if !os.IsNotExist(err) {
			return err
		}

		var result bytes.Buffer
		tmpl, err := template.New(f.path).Parse(f.content)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(&result, nil); err != nil {
			return err
		}

		err = ioutil.WriteFile(f.path, result.Bytes(), 0655)
		if err != nil {
			return err
		}
	}

	return nil
}
