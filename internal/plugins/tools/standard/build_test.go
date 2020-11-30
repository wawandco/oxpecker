package standard

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestBinaryOutput(t *testing.T) {
	c := &Plugin{}
	output := c.binaryOutput("aaa")

	if output != "bin/aaa" {
		t.Errorf("binaryOutput should be %v not %v", "bin/aaa", output)
	}
}

func TestComposeBuildArgs(t *testing.T) {
	td := t.TempDir()
	err := os.Chdir(td)
	if err != nil {
		t.Fatal(err)
	}

	file := `module wawandco/app`
	err = ioutil.WriteFile("go.mod", []byte(file), 0444)
	if err != nil {
		t.Fatal(err)
	}

	c := &Plugin{}
	c.static = true
	args, err := c.composeBuildArgs()
	if err != nil {
		t.Fatal("shoud not get error")
	}

	expectedArgs := []string{
		"build",
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,
		"-o",
		"bin/app",
		"./cmd/app",
	}

	if strings.Join(args, " ") != strings.Join(expectedArgs, " ") {
		t.Fatalf("args do not match should be %v,got %v", expectedArgs, args)
	}

	c.output = "other/bin"
	c.static = true
	args, err = c.composeBuildArgs()
	if err != nil {
		t.Fatal("shoud not get error")
	}

	expectedArgs = []string{
		"build",
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,
		"-o",
		"other/bin",
		"./cmd/app",
	}

	if strings.Join(args, " ") != strings.Join(expectedArgs, " ") {
		t.Fatalf("args do not match should be %v", expectedArgs)
	}

	c.buildTags = []string{"netdns=go"}

	args, err = c.composeBuildArgs()
	if err != nil {
		t.Fatal("shoud not get error")
	}

	expectedArgs = []string{
		"build",
		"--ldflags",
		"-linkmode external",
		"--ldflags",
		`-extldflags "-static"`,
		"-o",
		"other/bin",
		"-tags",
		"netdns=go",
		"./cmd/app",
	}

	if strings.Join(args, " ") != strings.Join(expectedArgs, " ") {
		t.Fatalf("args do not match should be %v, got %v", expectedArgs, args)
	}

	c.static = false
	args, err = c.composeBuildArgs()
	if err != nil {
		t.Fatal("shoud not get error")
	}
	expectedArgs = []string{
		"build",
		"-o",
		"other/bin",
		"-tags",
		"netdns=go",
		"./cmd/app",
	}

	if strings.Join(args, " ") != strings.Join(expectedArgs, " ") {
		t.Fatalf("args do not match should be %v,got %v", expectedArgs, args)
	}

}
