// Liquibase package aims to provide a ox plugin to manage migrations
// with the Liquibase style.  That is, using the databasechangelog and
// databasechangeloglock as well as the xml format for liquibase
// migrations.
//
// IMPORTANT: This plugin is not ready to be used in production, rather
// it aims to allow developers to avoid the java and liquibase installation
// by doing what liquibase does.
package liquibase

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/wawandco/oxpecker/plugins"
)

// Plugins on this package.
func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		&Command{connections: pop.Connections},
		&Generator{},
	}
}
