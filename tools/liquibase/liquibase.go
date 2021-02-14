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
	"fmt"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
	"github.com/wawandco/oxpecker/plugins"
)

// Plugins on this package.
func Plugins(conns interface{}) []plugins.Plugin {
	connections := map[string]URLProvider{}

	switch v := conns.(type) {
	case map[string]*pop4.Connection:
		for k, conn := range v {
			connections[k] = conn
		}
	case map[string]*pop5.Connection:
		for k, conn := range v {
			connections[k] = conn
		}
	default:
		fmt.Println("[warning] Liquibase plugin ONLY receives pop v4 and v5 connections")
	}

	return []plugins.Plugin{
		&Command{connections: connections},
		&Generator{},
	}
}
