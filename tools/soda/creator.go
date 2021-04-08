package soda

import (
	"fmt"

	"github.com/wawandco/oxpecker/tools/soda/fizz"
	"github.com/wawandco/oxpecker/tools/soda/sql"
)

var _ Creator = fizz.Creator{}
var _ Creator = sql.Creator{}

type Creator interface {
	// Receives a migration type string and returns if it
	// applies to it or not.
	Creates(string) bool

	// Creates the migration.
	Create(dir, name string, args []string) error
}

type Creators []Creator

func (c Creators) CreatorFor(name string) Creator {
	fmt.Println("CREATORS----------------------->", c)
	for _, x := range c {
		if x.Creates(name) {
			return x
		}
	}

	return nil
}
