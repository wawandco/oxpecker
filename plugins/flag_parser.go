package plugins

// FlagParser is a plugin that will use the params to parse flags
// that may affect the way it work.
type FlagParser interface {
	Plugin
	ParseFlags([]string) error
}
