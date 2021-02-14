package render

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestInitializer(t *testing.T) {
	t.Run("CompleteArgs", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join(root, "myapp"), 0777)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		var dx sync.Map
		dx.Store("module", "oosss/myapp")
		dx.Store("name", "myapp")
		dx.Store("folder", filepath.Join(root, "myapp"))

		err = i.Initialize(context.Background(), &dx)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "myapp", "app", "render", "render.go")
		bm, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal("should have created the file")
		}

		content := []string{
			`package render`,
			`var Engine = render.New(render.Options{`,
			`var Helpers = map[string]interface{}{`,
			`"partialFeeder": base.Templates.FindString,`,
		}

		for _, c := range content {
			if !bytes.Contains(bm, []byte(c)) {
				t.Errorf("`%v` does not contain `%v`", path, c)
			}
		}
	})

	t.Run("IncompleteArgs", func(t *testing.T) {
		root := t.TempDir()

		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join(root, "myapp"), 0777)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		var dx sync.Map

		err = i.Initialize(context.Background(), &dx)
		if err != ErrIncompleteArgs {
			t.Fatalf("error should be `%v`, got `%v`", ErrIncompleteArgs, err)
		}

		dx.Store("folder", filepath.Join(root, "myapp"))
		err = i.Initialize(context.Background(), &dx)
		if err != ErrIncompleteArgs {
			t.Fatalf("error should be `%v`, got `%v`", ErrIncompleteArgs, err)
		}

		dx.Store("module", "some/myapp")
		err = i.Initialize(context.Background(), &dx)
		if err != nil {
			t.Fatalf("error should be `%v`, got `%v`", nil, err)
		}
	})
}
