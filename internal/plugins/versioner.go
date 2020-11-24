package plugins

// Versioner interface allows a command to provide version
type Versioner interface {
	Plugin

	ToolName() string

	// Version returns the version or any error
	// when finding it.
	Version() (string, error)
}
