package build

// AfterBuilder interface allows to identify the things
// that will run after the build process has ended, things
// like cleanup and reverting go here
type AfterBuilder interface {
	Name() string
	AfterBuild(string, []string) error
}
