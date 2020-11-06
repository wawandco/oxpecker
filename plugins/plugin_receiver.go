package plugins

// PluginReceiver is an interface for those plugins that need to
// receive the list of plugins.
type PluginReceiver interface {
	Receive([]Plugin)
}
