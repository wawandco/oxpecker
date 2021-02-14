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
	app  *buffalo.App
)

// App creates a new application with default settings and reading 
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returing it
func New() *buffalo.App {
	if app != nil {
		return app
	}

	app = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_{{.Name}}_session",
	})

	// Setting the routes for the app
	setRoutes(app)

	return app
}
`

var routesGo = `

package app

import (
	"{{.Module}}/app/actions/home"
	"{{.Module}}/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(app *buffalo.App) {
	app.Use(middleware.Transaction)
	app.Use(middleware.ParameterLogger)
	app.Use(middleware.CSRF)

	app.GET("/", home.Home)
	app.ServeFiles("/", {{.Name}}.Assets)
}
`
