package standard

import (
	"github.com/spf13/pflag"
)

// ParseFlags
func (b *Plugin) ParseFlags(args []string) {
	b.flags = pflag.NewFlagSet(b.Name(), pflag.ContinueOnError)
	b.flags.StringVarP(&b.output, "output", "o", "", "the path the binary will be generated at")
	b.flags.StringSliceVarP(&b.buildTags, "tags", "", []string{}, "tags to pass the go build command")
	b.flags.BoolVar(&b.static, "static", true, `build a static binary using  --ldflags '-linkmode external -extldflags "-static"'`)
	b.flags.Parse(args)
}

// ParseFlags
func (b *Plugin) Flags() *pflag.FlagSet {
	return b.flags
}
