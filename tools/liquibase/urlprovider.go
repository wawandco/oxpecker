package liquibase

// URLProvider provider is a struct that has the URL()
// method. Like pop.Connection.
type URLProvider interface {
	URL() string
}
