package build

import (
	"github.com/paganotoni/oxpecker/internal/plugins"
	"github.com/spf13/pflag"
)

func (b *Command) ParseFlags(args []string) error {
	// TODO: This needs to happen with all of the plugins
	for _, plugin := range b.builders {
		fp, ok := plugin.(plugins.FlagParser)
		if !ok {
			continue
		}

		err := fp.ParseFlags(args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Command) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("build", pflag.ContinueOnError)
	for _, plugin := range b.buildPlugins {
		fp, ok := plugin.(plugins.FlagParser)
		if !ok {
			continue
		}

		fp.Flags().VisitAll(func(f *pflag.Flag) {
			fs.AddFlag(f)
		})
	}

	return fs
}
