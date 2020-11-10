package environment

import (
	"context"
	"fmt"
	"os"
)

func (g *GoEnv) RunBeforeTest(ctx context.Context, root string, args []string) error {
	fmt.Println("Running GO_ENV before test")
	return os.Setenv("GO_ENV", "test")
}
