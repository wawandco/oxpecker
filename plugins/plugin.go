package plugins

// Plugin interface is the base for the different plugins that can
// be attached to the plugin system. It is based on the `Name() string` method
// that will be useful for identification of the plugin.
//
// Other plugins (PluginReceivers) could specify other interfaces to identify
// plugins specific to what they do.
type Plugin interface {
	Name() string
}
