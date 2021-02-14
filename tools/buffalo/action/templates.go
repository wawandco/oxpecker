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

var actionsGo = `
package actions

import(
	"{{.Module}}/app/render"
)

var (
	// r is a buffalo/render Engine that will be used by actions 
	// on this package to render render HTML or any other formats.
	r = render.Engine
)
`

var actionsTestGo = `
package actions_test

import (
	"testing"
	"{{.Module}}/app"

	"github.com/gobuffalo/suite/v3"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	bapp, err := app.New()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	as := &ActionSuite{suite.NewAction(bapp)}
	suite.Run(t, as)
}
`

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
