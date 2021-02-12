package ox

var mainTemplate = `
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"{{.Module}}"
	_ "{{.Module}}/app/tasks"
	_ "{{.Module}}/app/models"


	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker"
	"github.com/wawandco/oxpecker/tools"
	"github.com/wawandco/oxpecker/tools/pop/migrate"
)

// main function for the tooling cli, will be invoked by Oxpecker
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
  	fmt.Print("[info] Using {{.Name}}/cmd/ox \n\n")
	ctx := context.Background()
    
  	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
    
	cl := cli.New()
	// append your plugins here
	cl.Plugins = append(cl.Plugins, tools.Base...)
	cl.Plugins = append(cl.Plugins, migrate.Plugins()...)
    
    err = cl.Run(ctx, pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
`
