package soda

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/log"
	"github.com/wawandco/ox/plugins"
)

var (
	// These are the interfaces we know that this
	// plugin must satisfy for its correct functionality
	_ plugins.Plugin     = (*Generator)(nil)
	_ plugins.FlagParser = (*Generator)(nil)
)

// Generator allows to identify model as a plugin
type Generator struct {
	creators Creators

	flags         *pflag.FlagSet
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

	cr := g.creators.CreatorFor(g.migrationType)
	if cr == nil {
		return errors.New("type not found")
	}

	dirPath := filepath.Join(root, "migrations")
	if !g.exists(dirPath) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}

	name := flect.Underscore(flect.Pluralize(args[2]))
	columns := g.parseColumns(args[3:])

	err := cr.Create(dirPath, name, columns)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) ParseFlags(args []string) {
	g.flags = pflag.NewFlagSet("type", pflag.ContinueOnError)
	g.flags.StringVarP(&g.migrationType, "type", "t", "fizz", "the type of the migration")
	g.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (g *Generator) Flags() *pflag.FlagSet {
	return g.flags
}

func (g *Generator) Receive(pls []plugins.Plugin) {
	for _, v := range pls {
		cr, ok := v.(Creator)
		if !ok {
			continue
		}

		g.creators = append(g.creators, cr)
	}
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func (g *Generator) parseColumns(args []string) []string {
	var columns []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			columns = append(columns, arg)
		}
	}

	return columns
}
