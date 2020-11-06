package packr

import (
	"context"

	"github.com/gobuffalo/packr/v2/jam"
)

// BeforeBuild pack
func (w Tool) BeforeBuild(ctx context.Context, root string, args []string) error {
	// Generate file on the root to have Packr packs files correctly.
	return nil
}

// Build uses the Packr Jam library to generate packd folders
// that contain those in the binary.
func (w Tool) Build(ctx context.Context, root string, args []string) error {
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
func (w Tool) AfterBuild(root string, args []string) error {
	return jam.Clean()
}
