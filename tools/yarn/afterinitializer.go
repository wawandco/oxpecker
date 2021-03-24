package yarn

import (
	"context"
	"os"
	"os/exec"
)

type AfterInitializer struct{}

func (ai AfterInitializer) Name() string {
	return "yarn/afterinitializer"
}

func (ai AfterInitializer) AfterInitialize(ctx context.Context) error {
	c := exec.CommandContext(ctx, "yarn", "install", "--no-progress")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}
