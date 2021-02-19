// Package tools contains plugins for different tools used in the
// development workflow.
package tools

import (
	"github.com/wawandco/oxpecker/lifecycle/build"
	"github.com/wawandco/oxpecker/lifecycle/dev"
	"github.com/wawandco/oxpecker/lifecycle/fix"
	"github.com/wawandco/oxpecker/lifecycle/generate"
	"github.com/wawandco/oxpecker/lifecycle/new"
	"github.com/wawandco/oxpecker/lifecycle/test"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxpecker/tools/buffalo/model"
	"github.com/wawandco/oxpecker/tools/buffalo/resource"
	"github.com/wawandco/oxpecker/tools/buffalo/template"
	"github.com/wawandco/oxpecker/tools/cli/help"
	"github.com/wawandco/oxpecker/tools/docker"
	"github.com/wawandco/oxpecker/tools/envy"
	"github.com/wawandco/oxpecker/tools/flect"
	"github.com/wawandco/oxpecker/tools/grift"
	"github.com/wawandco/oxpecker/tools/node"
	"github.com/wawandco/oxpecker/tools/ox"
	"github.com/wawandco/oxpecker/tools/refresh"
	"github.com/wawandco/oxpecker/tools/standard"
	"github.com/wawandco/oxpecker/tools/webpack"
	"github.com/wawandco/oxpecker/tools/yarn"
)

// Base plugins for applications lifecycle. While oxpecker
// has other plugins this list is the base that is used across most of
// the apps we do. Feel free to add the rest in your cmd/ox/main.go file.
var Base = []plugins.Plugin{
	&help.Command{},

	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&yarn.Plugin{},
	&envy.Developer{},

	// Application Lifecycle commands.
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
	&generate.Command{},
	&new.Command{},
	&grift.Command{},

	// Builders
	&node.Builder{},
	&standard.Builder{},

	// Fixers
	&standard.Fixer{},

	// Generators
	&ox.Generator{},
	&template.Generator{},
	&model.Generator{},
	&resource.Generator{},

	// Initializer
	&refresh.Initializer{},
	&flect.Initializer{},
	&docker.Initializer{},

	// Testers
	&standard.Tester{},
	&envy.Tester{},
}
