package ox

var mainTemplate = `
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/wawandco/oxpecker/cli"
)

func main() {
  	fmt.Print("~~~~ Using {{.Name}}/cmd/ox ~~~\n\n")
	ctx := context.Background()
    
  	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
  	}
    
	ox := cli.New()
	// append your plugins here
	// ox.Plugins = append(wawandco.Plugins, ...)
    
    err = ox.Run(ctx, pwd, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
`
