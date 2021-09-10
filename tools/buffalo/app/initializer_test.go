package app

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/lifecycle/new"
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
		options := new.Options{
			Name:   "myapp",
			Module: "oosss/myapp",
			Folder: filepath.Join(root, "myapp"),
		}

		err = i.Initialize(ctx, options)
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

}
