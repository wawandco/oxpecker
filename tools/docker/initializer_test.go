package docker

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/lifecycle/new"
)

func TestInitilizer(t *testing.T) {
	t.Run("dockerFileDoesNotExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}
		ctx := context.Background()
		options := new.Options{
			Folder: filepath.Join(root),
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		rootDoc := filepath.Join(root, ".dockerignore")
		rootFile := filepath.Join(root, "Dockerfile")

		_, err = os.Stat(rootDoc)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create .dockerignore file , %v", err)
		}
		_, err = os.Stat(rootFile)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create  Dockerfile file , %v", err)
		}
	})

}
