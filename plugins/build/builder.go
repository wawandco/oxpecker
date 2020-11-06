package build

import "context"

// Builder interface allows to set the build steps to be run.
type Builder interface {
	Name() string
	Build(context.Context, string, []string) error
}
