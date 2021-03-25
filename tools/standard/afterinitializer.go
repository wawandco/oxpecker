package standard

import (
	"context"
	"errors"
	"os"
	"os/exec"

	"github.com/wawandco/oxpecker/lifecycle/new"
)

var _ new.AfterInitializer = (*AfterInitializer)(nil)

type AfterInitializer struct{}

func (i AfterInitializer) Name() string {
	return "standard/afterinitializer"
}

// Initialize the go module
func (i *AfterInitializer) AfterInitialize(ctx context.Context) error {
	root := ctx.Value("folder")
	if root == nil {

		return errors.New("folder is needed")
	}

	err := os.Chdir(root.(string))
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
