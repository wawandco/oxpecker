package yarn

import (
	"context"
	"os"
	"os/exec"
)

// RunBeforeBuild attempts to run yarn install if it finds yarn.lock
func (p *Plugin) RunBeforeBuild(ctx context.Context, root string, args []string) error {
	cmd := p.buildCmd(ctx)
	if cmd == nil {
		return nil
	}

	return cmd.Run()
}

// build command will return the command if yarn.lock is found
// otherwise returns nil
func (p *Plugin) buildCmd(ctx context.Context) *exec.Cmd {
	_, err := os.Stat("yarn.lock")
	if os.IsNotExist(err) {
		return nil
	}

	c := exec.CommandContext(ctx, "yarn", "install", "--no-progress")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c
}
