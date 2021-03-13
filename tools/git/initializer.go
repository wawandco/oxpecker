package git

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

var (
	ErrIncompleteArgs = errors.New("incomplete args")
)

// Initializer
type Initializer struct{}

func (i Initializer) Name() string {
	return "model/initializer"
}

func (i *Initializer) Initialize(ctx context.Context) error {
	f := ctx.Value("folder")
	if f == nil {
		return ErrIncompleteArgs
	}

	folder := f.(string)
	keeps := []string{
		"migrations",
		"public",
	}

	for _, k := range keeps {
		err := os.MkdirAll(filepath.Join(folder, k), 0777)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join(folder, k, ".gitkeep"), []byte{}, 0777)
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
