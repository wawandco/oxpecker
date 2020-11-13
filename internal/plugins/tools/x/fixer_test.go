package x

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFix(t *testing.T) {
	t.Run("MainAndGoModExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		content = []byte("module github.com/some/cool/package")
		err = ioutil.WriteFile("go.mod", content, 0600)
		if err != nil {
			t.Fatalf("could not create go.mod file: %v", err)
		}
		f := Fixer{}
		err = f.Fix()
		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}

	})
	t.Run("MainMissing", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		f := Fixer{}
		err = f.Fix()
		if err != ErrFileMainNotExist {
			t.Fatalf("error should be nill, got %v", err)
		}

	})
	t.Run("GoModMissing", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		f := Fixer{}
		err = f.Fix()
		if err == nil {
			t.Fatalf("error should be %v, got nil", err)
		}

	})
}

func TestMoveFile(t *testing.T) {

	t.Run("MainExists", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		f := Fixer{}
		err = f.moveFile("julian")
		if err != nil {
			t.Fatal("movefile did not work")
		}

		dat, err := ioutil.ReadFile(filepath.Join("cmd", "julian", "main.go"))
		if err != nil {
			t.Fatal("movefile did not work (reading file)")
		}

		if !bytes.Contains(dat, content) {
			t.Fatal("did not move content correctly")
		}
	})

	t.Run("ModuleNameEmpty", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		f := Fixer{}
		err = f.moveFile("")
		if err != ErrModuleNameNeeded {
			t.Fatalf("needs to return ErrModuleNameNeeded, got %v", err)
		}
	})

	t.Run("MainMissing", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		f := Fixer{}
		err = f.moveFile("julian")
		if err == nil {
			t.Fatal("movefile did work")
		}
	})

	t.Run("MainMissing", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		err = os.MkdirAll("cmd", 0755)
		if err != nil {
			t.Fatal("did not create cmd folder")
		}

		f := Fixer{}
		err = f.moveFile("julian")
		if err != nil {
			t.Fatal("movefile did work")
		}

		_, err = os.Stat(filepath.Join("cmd", "julian", "main.go"))
		if err != nil {
			t.Fatalf("should have worked, got %v", err)
		}
	})
}

func TestFindModuleName(t *testing.T) {

	tcases := []struct {
		content      string
		nameExpected string
		errExpected  error
	}{
		{
			content:      "random module content",
			nameExpected: "",
			errExpected:  ErrModuleNameNotFound,
		},
		{
			content:      "module moduleFixer",
			nameExpected: "moduleFixer",
		},

		{
			content:      "module my/large/module/name",
			nameExpected: "name",
		},

		{
			content:      "module github.com/some/cool/package",
			nameExpected: "package",
		},
		// TO DO:
		{
			content: `//One with comment
					  module github.com/some/cool/comment`,
			nameExpected: "comment",
		},
		{
			content:      "",
			nameExpected: "",
			errExpected:  ErrModuleNameNotFound,
		},

		{
			content:      "// module name tricky in comment",
			nameExpected: "",
			errExpected:  ErrModuleNameNotFound,
		},
	}

	for _, tcase := range tcases {
		t.Run(tcase.content, func(t *testing.T) {
			err := os.Chdir(t.TempDir())
			if err != nil {
				t.Fatal("could not move to tmp folder")
			}

			content := []byte(tcase.content)
			err = ioutil.WriteFile("go.mod", content, 0600)
			if err != nil {
				t.Fatalf("could not create go.mod file: %v", err)
			}

			f := Fixer{}
			name, err := f.findModuleName()
			fmt.Println(name)
			if err != tcase.errExpected {
				t.Fatalf("error should be %v got %v", tcase.errExpected, err)
			}

			if name != tcase.nameExpected {
				t.Fatalf("module name should be %v got %v", tcase.nameExpected, name)
			}
		})

	}
}

func TestFileExists(t *testing.T) {

	t.Run("MainExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		content := []byte("package main // Generated!")
		err = ioutil.WriteFile("main.go", content, 0600)
		if err != nil {
			t.Fatal("could not create main file")
		}

		f := Fixer{}
		_, err = f.fileExists()
		if err != nil {
			t.Fatalf("Should found file, got %v", err)
		}
	})
	t.Run("MainNotExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		f := Fixer{}
		found, err := f.fileExists()
		if err != ErrFileMainNotExist {
			t.Fatalf("Should return %v, got %v", ErrFileMainNotExist, err)
		}

		if found {
			t.Fatalf("should have not found the file")
		}
	})
}
