package refresh

// Plugin takes care of refresh related operations that the CLI
// needs to take care. I contains methods to satisfy the needs
// of each of the commands that will call it.
type Plugin struct{}

func (w Plugin) Name() string {
	return "Refresh"
}
