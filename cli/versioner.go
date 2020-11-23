package cli

var (
	// The version of the CLI
	version = "0.0.1"
)

// Versioner plugin provides the version of the X tool.
type Versioner struct{}

func (v Versioner) Name() string {
	return "cli/versioner"
}

func (v Versioner) ToolName() string {
	return "ox"
}

func (v Versioner) Version() (string, error) {
	return version, nil
}
