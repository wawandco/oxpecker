package liquibase

import (
	"os"
	"os/exec"
)

func (lb Command) runUp() error {
	runArgs, err := lb.buildRunArgs()
	if err != nil {
		return err
	}

	runArgs = append(runArgs, []string{
		"--changeLogFile=./migrations/changelog.xml",
		"update",
	}...)

	c := exec.Command("liquibase", runArgs...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}
