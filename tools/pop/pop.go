package pop

import (
	"github.com/gobuffalo/packd"
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/tools/pop/migrate"
	"github.com/wawandco/oxplugins/tools/pop/migration"
)

func Plugins(migrations packd.Box) []plugins.Plugin {
	result := migrate.Plugins(migrations)
	result = append(result, &migration.Generator{})

	return result
}
