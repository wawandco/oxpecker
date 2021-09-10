package webpack

import (
	"context"
	"os"
	"os/exec"

	"github.com/wawandco/ox/internal/log"
)

func (w Plugin) Develop(ctx context.Context, root string) error {
	var cmd *exec.Cmd

	switch w.packageManagerType(root) {
	case javascriptPackageManagerYarn:
		cmd = exec.CommandContext(ctx, "yarn", "run", "dev")
	case javascriptPackageManagerNPM:
		cmd = exec.CommandContext(ctx, "npm", "run", "dev")
	case javascriptPackageManagerNone:
		log.Warn("did not find yarn.lock nor package-lock.json, skipping webpack build.")

		return nil
	}

	cmd.Env = append(
		os.Environ(),
		"NODE_ENV=development",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
