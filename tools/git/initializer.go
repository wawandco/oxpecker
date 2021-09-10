package git

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/lifecycle/new"
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	keeps := []string{
		"migrations",
		"public",
	}

	for _, k := range keeps {
		err := os.MkdirAll(filepath.Join(options.Folder, k), 0777)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join(options.Folder, k, ".gitkeep"), []byte{}, 0777)
		if err == nil {
			continue
		}

		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("buffalo-models-initializer", pflag.ContinueOnError)
}
