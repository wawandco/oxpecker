package source

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TODO: https://pkg.go.dev/golang.org/x/tools/imports

// Build a template and write it to the passed path in the filename
// if folder does not exist this function will take care of creating it,
// also if there is any error it will return that.
func Build(filename, source string, data interface{}) error {
	path := filepath.Dir(filename)
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filename).Parse(source)
	if err != nil {
		return err
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, data)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, sbf.Bytes(), 0777)
	if err != nil {
		return err
	}
	return nil
}
