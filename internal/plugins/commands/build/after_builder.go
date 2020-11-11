package build

import "github.com/paganotoni/x/internal/plugins"

// AfterBuilder interface allows to identify the things
// that will run after the build process has ended, things
// like cleanup and reverting go here
type AfterBuilder interface {
	plugins.Plugin
	RunAfterBuild(string, []string) error
}
