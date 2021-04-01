package main

import (
	"context"
	"os"

	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker/internal/log"
)

// This is the main oxpecker CLI that provides the basic functionality
// by using the base plugins.
func main() {
	err := cli.Wrap(context.Background(), os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
