package compiler

import "github.com/paganotoni/x/plugins"

var (
	// These are the interfaces we know that this
	// plugin must satisfy for its correct functionality
	_ plugins.Plugin     = (*Compiler)(nil)
	_ plugins.FlagParser = (*Compiler)(nil)
)

// Compiler takes care of compiling the go binary, this might be one of
// the last steps when we have done other phases of the build process.
type Compiler struct {
	output string
}

func (g Compiler) Name() string {
	return "compiler"
}
