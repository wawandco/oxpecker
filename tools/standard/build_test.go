package standard

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestBinaryOutput(t *testing.T) {
	c := &Builder{}
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

	c := &Builder{}
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

func Test_ParseFlags_Empty(t *testing.T) {
	c := &Builder{}
	c.ParseFlags([]string{})

	if c.output != "" {
		t.Errorf("output should be empty, was `%s`", c.output)
	}
}

func Test_ParseFlags_Value(t *testing.T) {
	c := &Builder{}
	c.ParseFlags([]string{"-o", "something"})

	expected := "something"
	if c.output != expected {
		t.Errorf("output should be `%s`, was `%s`", expected, c.output)
	}
}
