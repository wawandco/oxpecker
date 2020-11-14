package info

import (
	"fmt"
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

	t.Run("ModuleCases", func(t *testing.T) {
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

				name, err := BuildName()
				fmt.Println(name)
				if err != tcase.errExpected {
					t.Fatalf("error should be %v got %v", tcase.errExpected, err)
				}

				if name != tcase.nameExpected {
					t.Fatalf("module name should be %v got %v", tcase.nameExpected, name)
				}
			})

		}
	})

}
