package yarn

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/lifecycle/new"
)

type AfterInitializer struct{}

func (ai AfterInitializer) Name() string {
	return "yarn/afterinitializer"
}

func (ai AfterInitializer) AfterInitialize(ctx context.Context, options new.Options) error {
	c := exec.CommandContext(ctx, "yarn", "install", "--no-progress")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}
