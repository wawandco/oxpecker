package liquibase

import (
	"os"
	"os/exec"
	"strconv"
)

func (lb *Command) runDown() error {

	runArgs, err := lb.buildRunArgs()
	if err != nil {
		return err
	}

	if lb.steps == 0 {
		lb.steps = 1
	}

	runArgs = append(runArgs, []string{
		"--changeLogFile=./migrations/changelog.xml",
		"rollbackCount",
		strconv.Itoa(lb.steps),
	}...)

	c := exec.Command("liquibase", runArgs...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c.Run()
}
