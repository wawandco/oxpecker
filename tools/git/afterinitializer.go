package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type AfterInitializer struct{}

func (ri AfterInitializer) Name() string {
	return "git/repoinitializer"
}

func (ri AfterInitializer) AfterInitialize(ctx context.Context) error {
	_, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("[warning] Git repo was not initialized given git was not present")
		return nil
	}

	folder, ok := ctx.Value("folder").(string)
	if !ok {
		return errors.New("folder needed")
	}

	err = os.Chdir(folder)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
