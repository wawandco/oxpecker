package build

import "context"

// beforeBuilder interface allows to identify the things
// that will run before the build process has started.
type beforeBuilder interface {
	Name() string
	BeforeBuild(context.Context, string, []string) error
}
