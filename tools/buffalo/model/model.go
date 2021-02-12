package model

import (
	"strings"

	"github.com/gobuffalo/flect/name"
)

type opts struct {
	Name     name.Ident
	Original string
	Attrs    []attr
	Imports  []string
}

func (o opts) Char() string {
	return strings.ToLower(o.Original[:1])
}
