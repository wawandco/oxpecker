package refresh

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gobuffalo/here"
	"github.com/markbates/refresh/refresh"
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
	info, err := here.Dir(root)
	if err != nil {
		return &refresh.Configuration{}, err
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
		BinaryName:      path.Base(info.Module.Path) + "-build",
		BuildTargetPath: filepath.Join(root, "cmd", path.Base(info.Module.Path)),
	}

	return config, nil
}
