package compiler

import (
	"context"
	"os"
	"os/exec"

	"github.com/paganotoni/x/tools"
)

// Build runs the Go compiler to generate the desired binary. Assuming the
// Go executable installed and can be invoked with `go`.
//
// IMPORTANT: it uses the static build flags.
func (g Tool) Build(ctx context.Context, root string, args []string) error {
	name, err := tools.BuildName()
	if err != nil {
		return err
	}

	buildArgs := []string{
		"build",

		//--static
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,

		//-o
		"-o",
		"bin/" + name,

		"./cmd/" + name,
	}

	cmd := exec.CommandContext(ctx, "go", buildArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
