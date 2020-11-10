package info

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_BuildName_Success(t *testing.T) {
	td := t.TempDir()
	err := os.Chdir(td)
	if err != nil {
		t.Fatal(err)
	}

	file := `module wawandco/something`
	err = ioutil.WriteFile("go.mod", []byte(file), 0600)
	if err != nil {
		t.Fatal(err)
	}

	name, err := BuildName()
	if err != nil {
		t.Fail()
	}

	if name != "something" {
		t.Fail()
	}
}

func Test_BuildName_Failed(t *testing.T) {
	td := t.TempDir()
	err := os.Chdir(td)
	if err != nil {
		t.Fatal(err)
	}

	name, err := BuildName()
	if err == nil {
		t.Fail()
	}

	if name != "" {
		t.Fail()
	}
}
