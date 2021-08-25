package standard

import (
	"context"
	"os"
	"os/exec"
)

type GoModAfterGenerator struct{}

func (gag GoModAfterGenerator) Name() string {
	return "mod-tidy"
}

func (gag GoModAfterGenerator) AfterGenerate(context.Context, string, []string) error {
	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
