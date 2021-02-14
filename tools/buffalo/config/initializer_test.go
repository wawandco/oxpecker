package config

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

		bm, err := ioutil.ReadFile(filepath.Join(root, "myapp", "config", "database.yml"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bm, []byte(`{{ envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test?sslmode=disable" }}`)) {
			t.Fatalf("should contain %v envor expression", string(bm))
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

		dx.Store("name", "myapp")
		err = i.Initialize(context.Background(), &dx)
		if err != nil {
			t.Fatalf("error should be `%v`, got `%v`", nil, err)
		}
	})
}
