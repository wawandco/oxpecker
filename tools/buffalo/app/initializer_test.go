package app

import (
	"bytes"
	"context"
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
		ctx = context.WithValue(ctx, "module", "oosss/myapp")
		ctx = context.WithValue(ctx, "name", "myapp")
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))

		err = i.Initialize(ctx)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		bm, err := ioutil.ReadFile(filepath.Join(root, "myapp", "app", "app.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bm, []byte(`package app`)) {
			t.Fatal("should contain package name")
		}

		if !bytes.Contains(bm, []byte(`New() *buffalo.App {`)) {
			t.Fatal("should contain func signature")
		}

		bm, err = ioutil.ReadFile(filepath.Join(root, "myapp", "app", "routes.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bm, []byte(`package app`)) {
			t.Fatal("should contain package name")
		}

		if !bytes.Contains(bm, []byte(`func setRoutes(root *buffalo.App) {`)) {
			t.Fatal("should use contain func signature")
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

		ctx = context.WithValue(ctx, "module", "some/myapp")
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
