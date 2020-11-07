package compiler

import (
	"github.com/spf13/pflag"
)

// ParseFlags
func (b *Compiler) ParseFlags(args []string) error {
	flags := pflag.NewFlagSet(b.Name(), pflag.ContinueOnError)
	flags.StringVarP(&b.output, "output", "o", "models", "the path the model will be created in")

	return flags.Parse(args)
}
