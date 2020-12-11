package main

import (
	"context"
	"log"
	"os"

	plugins "github.com/wawandco/oxpecker-plugins"
	"github.com/wawandco/oxpecker/cli"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cli := cli.New()
	cli.Plugins = append(cli.Plugins, plugins.All...)

	err = cli.Wrap(context.Background(), pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
