package refresh

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/lifecycle/new"
)

var (
	// the filename we will use for the generated yml.
	filename = `.buffalo.dev.yml`

	ErrNameRequired   = errors.New("name argument is required")
	ErrIncompleteArgs = errors.New("incomplete args")
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "refresh/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {
	rootYML := filepath.Join(options.Folder, filename)

	config := refresh.Configuration{
		AppRoot:         ".",
		BuildTargetPath: "." + string(filepath.Separator) + filepath.Join(".", "cmd", options.Name),
		BuildPath:       "bin",
		BuildDelay:      200 * time.Nanosecond,
		BinaryName:      fmt.Sprintf("tmp-%v-build", options.Name),
		IgnoredFolders: []string{
			"vendor",
			"log",
			"logs",
			"assets",
			"public",
			"grifts",
			"tmp",
			"bin",
			"node_modules",
			".sass-cache",
		},

		IncludedExtensions: []string{".go", ".env"},
		EnableColors:       true,
		LogName:            "ox",
	}

	err := config.Dump(rootYML)
	if err != nil {
		return err
	}

	return nil
}

func (i *Initializer) ParseFlags([]string) {}
func (i *Initializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("refresh-initializer", pflag.ContinueOnError)
}
