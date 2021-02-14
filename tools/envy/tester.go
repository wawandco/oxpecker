package envy

import (
	"context"

	"github.com/gobuffalo/envy"
)

type Tester struct{}

func (t Tester) Name() string {
	return "envy/tester"
}

func (b *Tester) RunBeforeTest(ctx context.Context, root string, args []string) error {
	envy.Set("GO_ENV", "test")
	return nil
}
