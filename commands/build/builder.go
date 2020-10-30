package build

import "context"

// builder interface allows to set the build steps to be run.
type builder interface {
	Name() string
	Build(context.Context, string, []string) error
}
