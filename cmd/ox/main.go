package main

import (
	"context"
	"log"
	"os"

	"github.com/paganotoni/oxpecker/cli"
	"github.com/paganotoni/oxpecker/internal/plugins"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cli := cli.New()
	cli.Plugins = append(cli.Plugins, plugins.All...)

	err = cli.Run(context.Background(), pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
