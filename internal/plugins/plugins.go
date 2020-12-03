package plugins

import (
	"github.com/wawandco/oxpecker/internal/plugins/lifecycle/build"
	"github.com/wawandco/oxpecker/internal/plugins/lifecycle/dev"
	"github.com/wawandco/oxpecker/internal/plugins/lifecycle/fix"
	"github.com/wawandco/oxpecker/internal/plugins/lifecycle/test"
	"github.com/wawandco/oxpecker/internal/plugins/tools/packr"
	"github.com/wawandco/oxpecker/internal/plugins/tools/pop"
	"github.com/wawandco/oxpecker/internal/plugins/tools/pop/migrate"
	"github.com/wawandco/oxpecker/internal/plugins/tools/refresh"
	"github.com/wawandco/oxpecker/internal/plugins/tools/standard"
	"github.com/wawandco/oxpecker/internal/plugins/tools/webpack"
	"github.com/wawandco/oxpecker/internal/plugins/tools/yarn"
	"github.com/wawandco/oxpecker/plugins"
)

// All plugins in this package
var All = []plugins.Plugin{
	// IMPORTANT: order matters!
	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&packr.Plugin{},
	&pop.Plugin{},
	&migrate.Plugin{},
	&standard.Plugin{},
	&yarn.Plugin{},

	// Fixers
	&pop.Fixer{},
	&standard.Fixer{},

	// Developer Lifecycle plugins
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
}
