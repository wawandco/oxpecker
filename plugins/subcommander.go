package plugins

// Subcommander allows a plugin to say which are its subcommands.
type Subcommander interface {
	Command
	PluginReceiver

	Subcommands() []Command
}
