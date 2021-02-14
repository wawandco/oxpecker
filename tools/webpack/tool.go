package webpack

// Plugin takes care of webpack related operations that the CLI
// needs to take care. I contains methods to satisfy the needs
// of each of the commands that will call it.
type Plugin struct{}

func (w Plugin) Name() string {
	return "Webpack"
}
