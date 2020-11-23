package packr

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2/jam"
	"github.com/paganotoni/oxpecker/internal/info"
)

// BeforeBuild pack
func (w Plugin) RunBeforeBuild(ctx context.Context, root string, args []string) error {
	// check if there are Go files otherwise generate one
	// to prevent Packr from failing
	found := false
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Dir(path) != root || filepath.Ext(path) != ".go" {
			return nil
		}

		found = true
		return nil
	})

	if err != nil {
		return err
	}

	// If found any .go in the root folder it will
	// not do the rest of the things here.
	if found {
		return nil
	}

	name, err := info.BuildName()
	if err != nil {
		return err
	}

	filename := name + ".go"
	content := "package " + name

	return ioutil.WriteFile(filename, []byte(content), 0600)
}

// Build uses the Packr Jam library to generate packd folders
// that contain those in the binary.
func (w Plugin) Build(ctx context.Context, root string, args []string) error {
	err := jam.Clean(root)
	if err != nil {
		return err
	}

	err = jam.Pack(jam.PackOptions{
		Roots: []string{root},
	})

	return err
}

// AfterBuild runs the jam cleanup
func (w Plugin) RunAfterBuild(root string, args []string) error {
	return jam.Clean()
}
