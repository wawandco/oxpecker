package build

// afterBuilder interface allows to identify the things
// that will run after the build process has ended, things
// like cleanup and reverting go here
type afterBuilder interface {
	Name() string
	AfterBuild(string, []string) error
}
