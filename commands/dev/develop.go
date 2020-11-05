package dev

import "context"

// developer is a tool that will be invoked for development
// purposes. Things like webpack with --watch and refresh.
type developer interface {
	Name() string
	Develop(context.Context, string) error
}
