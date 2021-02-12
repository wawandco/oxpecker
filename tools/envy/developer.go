package envy

import (
	"context"
	"os"

	"github.com/gobuffalo/envy"
)

// Developer plugin sets the GO_ENV variable
// before app starts if not set.
type Developer struct{}

func (t Developer) Name() string {
	return "envy/developer"
}

func (b *Developer) BeforeDevelop(ctx context.Context, root string, args []string) error {
	if env := os.Getenv("GO_ENV"); env != "" {
		envy.Set("GO_ENV", env)

		return nil
	}

	os.Setenv("GO_ENV", "development")
	envy.Set("GO_ENV", "development")
	return nil
}
