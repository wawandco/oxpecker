package action

import (
	"html/template"

	"github.com/gobuffalo/flect"
)

var actionTemplate string = `package actions

func {{ properize .Name }}(c buffalo.Context) error {
	return nil
}
`
var actionTestTemplate string = `package actions

import (
	"testing"
)

func Test_{{ capitalize .Name }}(t *testing.T) {
	t.Fail("This test needs to be implemented!")
}`

var templateFuncs = template.FuncMap{
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
