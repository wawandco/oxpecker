package folders

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

var (
	ErrNameNeeded     = errors.New("app name must be specified")
	ErrFolderExists   = errors.New("folder exists")
	ErrIncompleteArgs = errors.New("did not receive expected data")

	// Folders that will be created when the initializer runs
	// the [name] parts will be replaced by the name of the app.
	folders = []string{
		filepath.Join("[name]"),
		filepath.Join("[name]", "public"),
		filepath.Join("[name]", "migrations"),
		filepath.Join("[name]", "config"),
		filepath.Join("[name]", "app"),
		filepath.Join("[name]", "app", "actions"),
		filepath.Join("[name]", "app", "middleware"),
		filepath.Join("[name]", "app", "assets"),
		filepath.Join("[name]", "app", "assets", "js"),
		filepath.Join("[name]", "app", "assets", "css"),
		filepath.Join("[name]", "app", "render"),
		filepath.Join("[name]", "app", "tasks"),
		filepath.Join("[name]", "cmd", "[name]"),
	}
)

// Initializer is in charge of building the bones of the
// Buffalo application. it will use the name argument and take the
// base part of it to build the folders.
// Some examples:
// - `ox new bongo` 					=> creates the new app in the bongo folder
// - `ox new github.com/wawandco/bongo` => creates the new app in the bongo folder
// - `ox new wawandco/bongo` 			=> creates the new app in the bongo folder
//
// This initialzier will return an error if the destination folder exists. The --force
// flag allows to remove and replace that folder.
type Initializer struct {
	// force folder creation if exists.
	force bool

	flags *pflag.FlagSet
}

// Name of the plugin
func (i Initializer) Name() string {
	return "folders/initializer"
}

// Initialize the app by creating the needed folders. It will infer the name of the
// folder from the args passed.
func (i *Initializer) Initialize(ctx context.Context) error {
	n := ctx.Value("name")
	if n == nil {
		return ErrNameNeeded
	}

	r := ctx.Value("root")
	if r == nil {
		return ErrIncompleteArgs
	}

	b := ctx.Value("folder")
	if b == nil {
		return ErrIncompleteArgs
	}

	root := r.(string)
	name := n.(string)
	base := b.(string)

	if _, err := os.Stat(base); err == nil && !i.force {
		return ErrFolderExists
	}

	err := os.RemoveAll(base)
	if err != nil {
		return err
	}

	for _, v := range folders {
		v = strings.ReplaceAll(v, "[name]", name)
		v = filepath.Join(root, v)

		err := os.MkdirAll(v, 0777)
		if err != nil {
			return err
		}

		fmt.Printf("[info] Created %v folder\n", v)
	}

	return nil
}

func (d *Initializer) ParseFlags(args []string) {
	d.flags = pflag.NewFlagSet(d.Name(), pflag.ContinueOnError)
	d.flags.BoolVarP(&d.force, "force", "f", false, "force the creation by removing folder if exists")
	d.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (d *Initializer) Flags() *pflag.FlagSet {
	return d.flags
}
