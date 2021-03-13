package standard

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/lifecycle/new"
)

var _ new.AfterInitializer = (*AfterInitializer)(nil)

type AfterInitializer struct{}

func (i AfterInitializer) Name() string {
	return "standard/afterinitializer"
}

// Initialize the go module
func (i *AfterInitializer) AfterInitialize(ctx context.Context, root string, args []string) error {
	fmt.Println(root)

	err := os.Chdir(root)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(context.Background(), "go", "mod", "tidy")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *AfterInitializer) ParseFlags(flags []string) {}

func (i *AfterInitializer) Flags() *pflag.FlagSet {
	return pflag.NewFlagSet("std/afterinit", pflag.ContinueOnError)
}
