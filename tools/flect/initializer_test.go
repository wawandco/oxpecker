package flect

import (
	"context"
	"os"
	"testing"

	"github.com/wawandco/ox/lifecycle/new"
)

func TestInitializer(t *testing.T) {
	t.Run("InflectionsFileDoesNotExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		ctx := context.Background()
		options := new.Options{
			Folder: root,
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		_, err = os.Stat(root)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create file ")
		}

	})
	t.Run("inflectionFileExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootYml := root + "/inflections.yml"
		_, err = os.Create(rootYml)
		if err != nil {
			t.Fatalf("Problem creating file, %v", err)
		}

		i := Initializer{}
		ctx := context.Background()
		options := new.Options{
			Folder: root,
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be type nil, got %v", err)
		}
	})
}
