package dev

import "context"

// Developer is a tool that will be invoked for development
// purposes. Things like webpack with --watch and refresh.
type Developer interface {
	Name() string
	Develop(context.Context, string) error
}
