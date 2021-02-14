package main

import (
	"context"
	"log"
	"os"

	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker/tools"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cli := cli.New()
	cli.Plugins = append(cli.Plugins, tools.Base...)

	err = cli.Wrap(context.Background(), pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
