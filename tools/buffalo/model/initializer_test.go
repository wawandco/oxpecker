package model

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
		ctx = context.WithValue(ctx, "module", "cool/myapp")
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))

		err = i.Initialize(ctx)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		_, err = os.Stat(filepath.Join(root, "myapp", "app", "models", "models_test.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		bmodels, err := ioutil.ReadFile(filepath.Join(root, "myapp", "app", "models", "models.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bmodels, []byte(`github.com/gobuffalo/pop/v5`)) {
			t.Fatal("models should contain pop import")
		}

		bmodelst, err := ioutil.ReadFile(filepath.Join(root, "myapp", "app", "models", "models_test.go"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bmodelst, []byte(`github.com/gobuffalo/suite/v3`)) {
			t.Fatal("models should contain suite import")
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

	})
}
