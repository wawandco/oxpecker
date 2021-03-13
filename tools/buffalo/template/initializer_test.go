package template

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
		ctx := context.Background()
		ctx = context.WithValue(ctx, "args", []string{"new", "oosss/myapp"})
		ctx = context.WithValue(ctx, "name", "myapp")
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))

		err = i.Initialize(ctx)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		bmodels, err := ioutil.ReadFile(filepath.Join(root, "myapp", "app", "templates", "application.plush.html"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		fmt.Println(string(bmodels))

		if !bytes.Contains(bmodels, []byte(`<!DOCTYPE html>`)) {
			t.Fatal("should contain HTML heading")
		}

		if !bytes.Contains(bmodels, []byte(`<%= javascriptTag("application.js", {defer: true}) %>`)) {
			t.Fatal("do the js deferred thing")
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
		ctx := context.Background()
		err = i.Initialize(ctx)
		if err != ErrIncompleteArgs {
			t.Fatalf("error should be `%v`, got `%v`", ErrIncompleteArgs, err)
		}

		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))
		err = i.Initialize(ctx)
		if err != ErrIncompleteArgs {
			t.Fatalf("error should be `%v`, got `%v`", ErrIncompleteArgs, err)
		}

		ctx = context.WithValue(ctx, "name", "myapp")
		err = i.Initialize(ctx)
		if err != nil {
			t.Fatalf("error should be `%v`, got `%v`", nil, err)
		}

	})
}
