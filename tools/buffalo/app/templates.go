package app

var appGo = `
package app

import (
	"{{.Module}}/app/models"
	"{{.Module}}/app/render"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

var (
	root  *buffalo.App
)

// App creates a new application with default settings and reading 
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returing it
func New() *buffalo.App {
	if root != nil {
		return app
	}

	root = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_{{.Name}}_session",
	})

	// Setting the routes for the app
	setRoutes(root)

	return root
}
`

var routesGo = `

package app

import (
	base "{{.Module}}"
	"{{.Module}}/app/actions/home"
	"{{.Module}}/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Show)
	root.ServeFiles("/", base.Assets)
}
`
