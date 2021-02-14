package refresh

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/pflag"
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

func (i *Initializer) Initialize(ctx context.Context, dx *sync.Map) error {
	n, ok := dx.Load("name")
	if !ok {
		return ErrNameRequired
	}

	folder, ok := dx.Load("folder")
	if !ok {
		return ErrNameRequired
	}

	rootYML := filepath.Join(folder.(string), filename)
	name := n.(string)

	config := refresh.Configuration{
		AppRoot:         ".",
		BuildTargetPath: filepath.Join(".", "cmd", name),
		BuildPath:       "bin",
		BuildDelay:      200 * time.Nanosecond,
		BinaryName:      fmt.Sprintf("tmp-%v-build", name),
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
