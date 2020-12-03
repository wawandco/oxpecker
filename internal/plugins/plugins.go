package plugins

import (
	"github.com/paganotoni/oxpecker/internal/plugins/lifecycle/build"
	"github.com/paganotoni/oxpecker/internal/plugins/lifecycle/dev"
	"github.com/paganotoni/oxpecker/internal/plugins/lifecycle/fix"
	"github.com/paganotoni/oxpecker/internal/plugins/lifecycle/test"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/packr"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/pop"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/pop/migrate"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/refresh"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/standard"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/webpack"
	"github.com/paganotoni/oxpecker/internal/plugins/tools/yarn"
	"github.com/paganotoni/oxpecker/plugins"
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

	// Developer Lifecycle plugins
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
}
