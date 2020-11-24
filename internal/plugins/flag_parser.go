package plugins

import "github.com/spf13/pflag"

// FlagParser is a plugin that will use the params to parse flags
// that may affect the way it work.
type FlagParser interface {
	Plugin

	ParseFlags([]string) error
	Flags() *pflag.FlagSet
}
