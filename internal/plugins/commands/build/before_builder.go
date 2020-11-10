package build

import "context"

// BeforeBuilder interface allows to identify the things
// that will run before the build process has started.
type BeforeBuilder interface {
	Name() string
	BeforeBuild(context.Context, string, []string) error
}
