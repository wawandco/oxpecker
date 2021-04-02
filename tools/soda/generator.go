package soda

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/internal/log"
	"github.com/wawandco/oxpecker/tools/soda/creator"
)

// Generator allows to identify model as a plugin
type Generator struct {
	migrationType string
}

// Name returns the name of the plugin
func (g Generator) Name() string {
	return "pop/generate-migration"
}

// InvocationName is used to identify the generator when
// the generate command is called.
func (g Generator) InvocationName() string {
	return "migration"
}

// Generate generates an empty [name].plush.html file
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		log.Info("No name specified, please use `ox generate migration [name] [columns?] --type=[sql|fizz]`")
		return nil
	}

	g.parseFlag(args)

	dirPath := filepath.Join(root, "migrations")
	if !g.exists(dirPath) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}

	creator, err := creator.CreateMigrationFor(strings.ToLower(g.migrationType))
	if err != nil {
		return err
	}

	name := flect.Underscore(flect.Pluralize(strings.ToLower(args[2])))
	columns := g.parseColumns(args[2:])

	if err = creator.Create(dirPath, columns); err != nil {
		return errors.Wrap(err, "failed creating migrations")
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s", timestamp, name)

	log.Infof("generated: migrations/%s.up.%s", fileName, creator.Name())
	log.Infof("generated: migrations/%s.down.%s", fileName, creator.Name())

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func (g *Generator) parseFlag(args []string) {
	flags := pflag.NewFlagSet("type", pflag.ContinueOnError)
	flags.StringVarP(&g.migrationType, "type", "t", "fizz", "the type of the migration")
	flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (g *Generator) parseColumns(args []string) []string {
	if len(args) == 1 {
		return args
	}

	var columns []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			columns = append(columns, arg)
		}
	}

	return columns
}
