package main

import (
	"context"
	"os"

	"github.com/wawandco/ox/cli"
	"github.com/wawandco/ox/internal/log"
)

// This is the main ox CLI that provides the basic functionality
// by using the base plugins.
func main() {
	err := cli.Wrap(context.Background(), os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
