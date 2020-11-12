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

	cli := cli.NewWithRoot(pwd)
	err = cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
