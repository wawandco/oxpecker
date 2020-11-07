package compiler

import (
	"fmt"

	"github.com/paganotoni/x/plugins"
	"github.com/spf13/pflag"
)

var _ plugins.FlagParser = (*Tool)(nil)

// ParseFlags
func (b *Tool) ParseFlags(args []string) error {
	fmt.Println(">> Parsing args on build ")
	flags := pflag.NewFlagSet(b.Name(), pflag.ContinueOnError)

	flags.StringVarP(&b.output, "output", "o", "models", "the path the model will be created in")
	// flags.StringVarP(&g.modelPkg, "pkg", "", "models", "the import part the model will be created in")
	// flags.StringVarP(&g.structTag, "struct-tag", "", "json", "sets the struct tags for model (xml/json/jsonapi)")

	return flags.Parse(args)
}
