package plugins

// HelpTexter interface allows plugins to provide text to the
// help command or subcommand.
type HelpTexter interface {
	// Help returns a string describing what the command/subcommand does.
	HelpText() string
}
