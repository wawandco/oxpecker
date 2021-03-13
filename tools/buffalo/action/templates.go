package action

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
	"{{.}}/app/render"
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
	"{{.}}/app"

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

var homeGo = `
package home

import (
	"{{.}}/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Show(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home.plush.html"))
}
`
