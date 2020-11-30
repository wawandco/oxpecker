package standard

import (
	"context"
	"os"
	"os/exec"

	"github.com/paganotoni/oxpecker/internal/info"
)

// Build runs the Go compiler to generate the desired binary. Assuming the
// Go executable installed and can be invoked with `go`.
//
// IMPORTANT: it uses the static build flags.
func (g *Plugin) Build(ctx context.Context, root string, args []string) error {
	buildArgs, err := g.composeBuildArgs()
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "go", buildArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func (g *Plugin) composeBuildArgs() ([]string, error) {
	name, err := info.BuildName()
	if err != nil {
		return []string{}, err
	}

	buildArgs := []string{
		"build",
	}
	//static
	static := []string{
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,
	}
	if g.static {
		for _, st := range static {
			buildArgs = append(buildArgs, st)
		}
	}
	//o
	o := []string{
		"-o",
		g.binaryOutput(name),
	}
	for _, oline := range o {
		buildArgs = append(buildArgs, oline)
	}

	// add the build

	if len(g.buildTags) != 0 {
		buildArgs = append(buildArgs, "-tags")
		for _, tag := range g.buildTags {
			buildArgs = append(buildArgs, tag)
		}
	}

	buildArgs = append(buildArgs, "./cmd/"+name)

	return buildArgs, nil
}

// binaryOutput considers the output passed to
// use it or default to bin/name.
func (g *Plugin) binaryOutput(name string) string {
	output := "bin/" + name
	if g.output != "" {
		output = g.output
	}

	return output
}
