package dev

import "context"

type (
	// Developer is a tool that will be invoked for development
	// purposes. Things like webpack with --watch and refresh.
	Developer interface {
		Name() string
		Develop(context.Context, string) error
	}

	// BeforeDeveloper is an interface that will allow to identify those plugins that
	// need to run before the dev commands is run.
	BeforeDeveloper interface {
		Name() string
		BeforeDevelop(context.Context, string) error
	}
)
