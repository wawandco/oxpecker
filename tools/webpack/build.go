package webpack

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wawandco/ox/internal/log"
)

const (
	javascriptPackageManagerYarn = "YARN"
	javascriptPackageManagerNPM  = "NPM"
	javascriptPackageManagerNone = "NONE"
)

// packageManagerType returns yarn or npm depending on the files
// found in the folder where the CLI is run. The first file found
// will determine the Javascript package manager.
func (w Plugin) packageManagerType(root string) string {
	info, err := os.Stat(filepath.Join(root, "yarn.lock"))
	if err == nil && !info.IsDir() {
		return javascriptPackageManagerYarn
	}

	info, err = os.Stat(filepath.Join(root, "package-lock.json"))
	if err == nil && !info.IsDir() {
		return javascriptPackageManagerNPM
	}

	return javascriptPackageManagerNone
}

// Build runs webpack build from the package.json scripts.
// if the project uses yarn will run `yarn run build`,
// if the project uses npm will run `npm run build`.
// otherwise will not run any of those.
//
// [Important] it assumes:
// - that there is a build script in package.json.
// - that yarn or npm is installed in the system.
func (w Plugin) Build(ctx context.Context, root string, args []string) error {
	var cmd *exec.Cmd
	switch w.packageManagerType(root) {
	case javascriptPackageManagerYarn:
		cmd = exec.CommandContext(ctx, "yarn", "run", "build")
	case javascriptPackageManagerNPM:
		cmd = exec.CommandContext(ctx, "npm", "run", "build")
	case javascriptPackageManagerNone:
		log.Warn("did not find yarn.lock nor package-lock.json, skipping webpack build.")

		return nil
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
