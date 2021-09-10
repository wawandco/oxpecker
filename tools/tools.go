// Package tool`s` contains plugins for different tools used in the
// development workflow.
package tools

import (
	"github.com/wawandco/ox/lifecycle/build"
	"github.com/wawandco/ox/lifecycle/dev"
	"github.com/wawandco/ox/lifecycle/fix"
	"github.com/wawandco/ox/lifecycle/generate"
	"github.com/wawandco/ox/lifecycle/new"
	"github.com/wawandco/ox/lifecycle/test"
	"github.com/wawandco/ox/plugins"
	"github.com/wawandco/ox/tools/buffalo/action"
	"github.com/wawandco/ox/tools/buffalo/app"
	"github.com/wawandco/ox/tools/buffalo/assets"
	"github.com/wawandco/ox/tools/buffalo/cmd"
	"github.com/wawandco/ox/tools/buffalo/config"
	"github.com/wawandco/ox/tools/buffalo/embedded"
	"github.com/wawandco/ox/tools/buffalo/middleware"
	"github.com/wawandco/ox/tools/buffalo/model"
	"github.com/wawandco/ox/tools/buffalo/render"
	"github.com/wawandco/ox/tools/buffalo/resource"
	"github.com/wawandco/ox/tools/buffalo/template"
	"github.com/wawandco/ox/tools/cli/help"
	"github.com/wawandco/ox/tools/cli/version"
	"github.com/wawandco/ox/tools/db"
	"github.com/wawandco/ox/tools/docker"
	"github.com/wawandco/ox/tools/envy"
	"github.com/wawandco/ox/tools/flect"
	"github.com/wawandco/ox/tools/git"
	"github.com/wawandco/ox/tools/grift"
	"github.com/wawandco/ox/tools/node"
	"github.com/wawandco/ox/tools/ox"
	"github.com/wawandco/ox/tools/refresh"
	"github.com/wawandco/ox/tools/soda"
	"github.com/wawandco/ox/tools/soda/fizz"
	"github.com/wawandco/ox/tools/soda/sql"
	"github.com/wawandco/ox/tools/standard"
	"github.com/wawandco/ox/tools/webpack"
	"github.com/wawandco/ox/tools/yarn"
)

// Base plugins for applications lifecycle. While ox
// has other plugins this list is the base that is used across most of
// the apps we build and maintain.
var Base = []plugins.Plugin{
	&help.Command{},

	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},
	&envy.Developer{},
	&db.CreateCommand{},
	&db.DropCommand{},
	&db.ResetCommand{},

	// Application Lifecycle commands.
	&build.Command{},
	&dev.Command{},
	&db.Command{},
	&test.Command{},
	&fix.Command{},
	&generate.Command{},
	&new.Command{},
	&grift.Command{},
	&version.Command{},

	// Builders
	&node.Builder{},
	&standard.Builder{},

	// Fixers
	&standard.Fixer{},

	// Generators
	&ox.Generator{},
	&template.Generator{},
	&model.Generator{},
	&action.Generator{},
	&resource.Generator{},
	&grift.Generator{},
	&soda.Generator{},

	// Initializer
	&embedded.Initializer{},
	&model.Initializer{},
	&render.Initializer{},
	&refresh.Initializer{},
	&template.Initializer{},
	&flect.Initializer{},
	&docker.Initializer{},
	&action.Initializer{},
	&middleware.Initializer{},
	&cmd.Initializer{},
	&config.Initializer{},
	&docker.Initializer{},
	&app.Initializer{},
	&standard.Initializer{},
	&grift.Initializer{},
	&assets.Initializer{},
	&soda.Initializer{},
	&git.Initializer{},

	&standard.AfterInitializer{},
	&yarn.AfterInitializer{},
	&git.AfterInitializer{},

	// Testers
	&standard.Tester{},
	&envy.Tester{},

	// Migration Creators
	&fizz.Creator{},
	&sql.Creator{},

	// Aftergenerators
	&standard.GoModAfterGenerator{},
}
