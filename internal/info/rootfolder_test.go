package info

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_RootFolder(t *testing.T) {
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

	t.Run("CurrentDir", func(t *testing.T) {
		folder := RootFolder()
		wd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		if folder != wd {
			t.Fatalf("Should return `%v` and got `%v`", wd, folder)
		}
	})

	t.Run("InnerFolder", func(t *testing.T) {
		wd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		inner := filepath.Join(wd, "inner", "folder")

		err = os.MkdirAll(inner, 0777)
		if err != nil {
			t.Fatal(err)
		}

		err = os.Chdir(inner)
		if err != nil {
			t.Fatal(err)
		}

		folder := RootFolder()
		if folder != wd {
			t.Fatalf("Should return `%v` and got `%v`", td, folder)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		td := t.TempDir()
		err := os.Chdir(td)
		if err != nil {
			t.Fatal(err)
		}

		folder := RootFolder()
		if folder != "" {
			t.Fatalf("Should return `%v` and got `%v`", td, folder)
		}
	})
}
