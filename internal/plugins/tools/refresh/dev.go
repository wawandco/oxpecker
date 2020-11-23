package refresh

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/markbates/refresh/refresh"
	"github.com/paganotoni/oxpecker/internal/info"
)

func (w Plugin) Develop(ctx context.Context, root string) error {
	config, err := w.config(root)
	if err != nil {
		return err
	}

	r := refresh.NewWithContext(config, ctx)
	return r.Start()
}

// config tries to pull the config from .buffalo.dev.yml otherwise uses default config
func (w Plugin) config(root string) (*refresh.Configuration, error) {
	c := &refresh.Configuration{}
	if _, err := os.Stat(".buffalo.dev.yml"); err == nil {
		err = c.Load(".buffalo.dev.yml")
		return c, err
	}

	return w.defaultConfig(root)
}

func (w Plugin) defaultConfig(root string) (*refresh.Configuration, error) {
	name, err := info.BuildName()
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
