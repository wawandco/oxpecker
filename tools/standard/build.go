package standard

import (
	"context"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
	"github.com/wawandco/ox/internal/info"
	"github.com/wawandco/ox/plugins"
)

var (
	// These are the interfaces we know that this
	// plugin must satisfy for its correct functionality
	_ plugins.Plugin     = (*Builder)(nil)
	_ plugins.FlagParser = (*Builder)(nil)
)

type Builder struct {
	output    string
	buildTags []string
	static    bool
	flags     *pflag.FlagSet
}

func (b Builder) Name() string {
	return "standard/builder"
}

// Build runs the Go compiler to generate the desired binary. Assuming the
// Go executable installed and can be invoked with `go`.
func (g *Builder) Build(ctx context.Context, root string, args []string) error {
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

// ParseFlags
func (b *Builder) ParseFlags(args []string) {
	b.flags = pflag.NewFlagSet(b.Name(), pflag.ContinueOnError)
	b.flags.StringVarP(&b.output, "output", "o", "", "the path the binary will be generated at")
	b.flags.StringSliceVarP(&b.buildTags, "tags", "", []string{}, "tags to pass the go build command")
	b.flags.BoolVar(&b.static, "static", true, `build a static binary using  --ldflags '-linkmode external -extldflags "-static"'`)
	b.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

// ParseFlags
func (b *Builder) Flags() *pflag.FlagSet {
	return b.flags
}

func (g *Builder) composeBuildArgs() ([]string, error) {
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
		buildArgs = append(buildArgs, static...)
	}

	// output
	o := []string{
		"-o",
		g.binaryOutput(name),
	}
	buildArgs = append(buildArgs, o...)

	// add the build tags
	if len(g.buildTags) != 0 {
		buildArgs = append(buildArgs, "-tags")
		buildArgs = append(buildArgs, g.buildTags...)
	}

	buildArgs = append(buildArgs, "./cmd/"+name)

	return buildArgs, nil
}

// binaryOutput considers the output passed to
// use it or default to bin/name.
func (g *Builder) binaryOutput(name string) string {
	output := "bin/" + name
	if g.output != "" {
		output = g.output
	}

	return output
}
