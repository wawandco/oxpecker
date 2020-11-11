package packr

// Plugin provides operations with Plugin for the CLI.
type Plugin struct{}

func (w Plugin) Name() string {
	return "Packr"
}
