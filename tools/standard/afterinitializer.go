package standard

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/lifecycle/new"
)

var _ new.AfterInitializer = (*AfterInitializer)(nil)

type AfterInitializer struct{}

func (i AfterInitializer) Name() string {
	return "standard/afterinitializer"
}

// Initialize the go module
func (i *AfterInitializer) AfterInitialize(ctx context.Context, options new.Options) error {
	err := os.Chdir(options.Folder)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
