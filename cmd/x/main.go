package main

import (
	"log"
	"os"

	"github.com/paganotoni/x/cli"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("GO111MODULE", "on") // Modules must be ON
	os.Setenv("CGO_ENABLED", "0")  // CGO disabled

	cli := cli.NewWithRoot(pwd)
	err = cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
