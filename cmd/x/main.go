package main

import (
	"log"
	"os"

	"github.com/paganotoni/x/cli"
)

func main() {
	cli := cli.NewCLI()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Modules must be ON
	os.Setenv("GO111MODULE", "on")

	// CGO disabled
	os.Setenv("CGO_ENABLED", "0")

	err = cli.Run(pwd, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
