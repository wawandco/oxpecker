package refresh

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/markbates/refresh/refresh"
	"golang.org/x/mod/modfile"
)

func (w Tool) Develop(ctx context.Context, root string) error {
	config, err := w.config(root)
	if err != nil {
		return err
	}

	r := refresh.NewWithContext(config, ctx)
	return r.Start()
}

// config tries to pull the config from .buffalo.dev.yml otherwise uses default config
func (w Tool) config(root string) (*refresh.Configuration, error) {
	c := &refresh.Configuration{}
	if _, err := os.Stat(".buffalo.dev.yml"); err == nil {
		err = c.Load(".buffalo.dev.yml")
		return c, err
	}

	return w.defaultConfig(root)
}

func (w Tool) defaultConfig(root string) (*refresh.Configuration, error) {
	name, err := buildName()
	if err != nil {
		return nil, err
	}

	config := &refresh.Configuration{
		IgnoredFolders: []string{
			"vendor",
			"log",
			"logs",
			"webpack",
			"public",
			"grifts",
			"tmp",
			"bin",
			"node_modules",
			".sass-cache",
		},

		IncludedExtensions: []string{
			".go",
			".mod",
			".env",
		},

		BuildPath:    "tmp",
		BuildDelay:   200 * time.Millisecond,
		EnableColors: true,
		LogName:      "x",

		// BuildFlags:   bflags,

		AppRoot:         root,
		BinaryName:      name + "-build",
		BuildTargetPath: filepath.Join(root, "cmd", name),
	}

	return config, nil
}

// buildName extracts the last part of the module by splitting on `/`
// this last part is useful for name of the binary and other things.
func buildName() (string, error) {
	content, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	path := modfile.ModulePath(content)
	name := filepath.Base(path)

	return name, nil
}
