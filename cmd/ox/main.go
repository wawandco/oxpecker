package main

import (
	"context"
	"os"

	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker/internal/log"
	"github.com/wawandco/oxpecker/tools"
)

func main() {
	cli := cli.New()
	cli.Plugins = append(cli.Plugins, tools.Base...)

	err := cli.Wrap(context.Background(), os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
