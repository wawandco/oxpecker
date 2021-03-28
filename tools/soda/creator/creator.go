package creator

// Creator follows a set of steps to make possible the creation
// of different migration types for example:
// -- fizz
// -- sql
// -- liquibase
type Creator interface {
	Name() string
	Create(string, []string) error
}
