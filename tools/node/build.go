package node

import (
	"context"
	"os"
)

// Builder sets up the NODE_ENV variable
type Builder struct{}

func (b Builder) Name() string {
	return "node/build"
}

func (b Builder) RunBeforeBuild(ctx context.Context, root string, args []string) error {
	return os.Setenv("NODE_ENV", os.Getenv("GO_ENV"))
}
