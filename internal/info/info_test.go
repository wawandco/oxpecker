package info

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBuildName(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		td := t.TempDir()
		err := os.Chdir(td)
		if err != nil {
			t.Fatal(err)
		}

		file := `module wawandco/something`
		err = ioutil.WriteFile("go.mod", []byte(file), 0444)
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
	})

	t.Run("Failed", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal(err)
		}

		err = os.Remove("go.mod")
		if err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}

		name, err := BuildName()
		if err == nil {
			t.Fail()
		}

		if name != "" {
			t.Fail()
		}
	})

}
