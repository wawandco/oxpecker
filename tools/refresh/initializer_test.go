package refresh

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInitializer(t *testing.T) {
	t.Run("Empty directory", func(t *testing.T) {
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
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "myapp")
		ctx = context.WithValue(ctx, "args", []string{"new", "cool/myapp"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))

		err = i.Initialize(ctx)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "myapp", ".buffalo.dev.yml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Fatalf("Did not create file in %v", path)
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal("could not read the file")
		}

		if !bytes.Contains(d, []byte("myapp")) {
			t.Fatal("did not containt app name")
		}

	})
	t.Run("BuffaloFileExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join(root, "myapp"), 0777)
		if err != nil {
			t.Fatal("could not create dir")
		}

		rootYml := filepath.Join(root, "myapp", ".buffalo.dev.yml")
		_, err = os.Create(rootYml)
		if err != nil {
			t.Fatalf("Problem creating file, %v", err)
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "myapp")
		ctx = context.WithValue(ctx, "args", []string{"new", "cool/myapp"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "myapp"))

		i := Initializer{}

		err = i.Initialize(ctx)

		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

	})
}
