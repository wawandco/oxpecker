package template

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

		bmodels, err := ioutil.ReadFile(filepath.Join(root, "myapp", "app", "templates", "application.plush.html"))
		if err != nil {
			t.Fatal("should have created the file")
		}

		if !bytes.Contains(bmodels, []byte(`<!DOCTYPE html>`)) {
			t.Fatal("should contain HTML heading")
		}

		if !bytes.Contains(bmodels, []byte(`<%= javascriptTag("application.js", {defer: true}) %>`)) {
			t.Fatal("do the js deferred thing")
		}
	})
}
