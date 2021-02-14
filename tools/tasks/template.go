package tasks

import (
	"html/template"

	"github.com/gobuffalo/flect"
)

var taskTemplate string = `package tasks

var _ = grift.Namespace("{{ .Name }}", func() error{
	return nil
})
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
