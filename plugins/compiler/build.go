package compiler

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/mod/modfile"
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

	output := "bin/" + name
	if g.output != "" {
		output = g.output
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
		output,

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
	content, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	path := modfile.ModulePath(content)
	name := filepath.Base(path)

	return name, nil
}
