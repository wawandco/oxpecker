package compiler

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/here"
)

// Build runs the Go compiler to generate the desired binary. Assuming the
// Go executable installed and can be invoked with `go`.
//
// IMPORTANT: it uses the static build flags.
func (g Tool) Build(ctx context.Context, root string, args []string) error {
	name, err := buildName()

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

// buildName extracts the last part of the module by splitting on `/`
// this last part is useful for name of the binary and other things.
func buildName() (string, error) {
	info, err := here.Current()
	if err != nil {
		return "", err
	}

	parts := strings.Split(info.Module.Path, "/")
	name := parts[len(parts)-1]

	return name, nil
}
