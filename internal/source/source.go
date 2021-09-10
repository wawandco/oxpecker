package source

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
	"github.com/wawandco/ox/internal/log"
)

var (
	helpers = template.FuncMap{
		"capitalize": func(field string) string {
			return flect.Capitalize(field)
		},
		"pascalize": func(field string) string {
			return flect.Pascalize(field)
		},
		"pluralize": func(field string) string {
			return flect.Pluralize(flect.Capitalize(field))
		},
		"properize": func(field string) string {
			return flect.Capitalize(flect.Singularize(field))
		},
		"singularize": func(field string) string {
			return flect.Singularize(field)
		},
		"underscore": func(field string) string {
			return flect.Underscore(field)
		},
	}
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

	tmpl := template.New(filename).Funcs(helpers)
	tmpl, err = tmpl.Parse(source)
	if err != nil {
		return errors.Wrap(err, "error intializing template")
	}

	sbf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sbf, data)
	if err != nil {
		return errors.Wrap(err, "error executing template")
	}

	err = ioutil.WriteFile(filename, sbf.Bytes(), 0777)
	if err != nil {
		return errors.Wrap(err, "error writing file")
	}

	log.Infof("generated %v", filename)
	return nil
}
