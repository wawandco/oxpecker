package webpack

// Tool takes care of webpack related operations that the CLI
// needs to take care. I contains methods to satisfy the needs
// of each of the commands that will call it.
type Tool struct{}

func (w Tool) Name() string {
	return "Webpack"
}
