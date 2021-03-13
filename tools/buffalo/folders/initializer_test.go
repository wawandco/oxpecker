package folders

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitializerRun(t *testing.T) {

	t.Run("valid args", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "app")
		ctx = context.WithValue(ctx, "args", []string{"new", "app"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(ctx)
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}

		for _, v := range folders {
			v = strings.ReplaceAll(v, "[name]", "app")
			v = filepath.Join(root, v)

			if _, err := os.Stat(v); os.IsNotExist(err) {
				t.Errorf("should have created %v", v)
			}
		}
	})

	t.Run("invalid args", func(t *testing.T) {
		i := Initializer{}
		ctx := context.Background()
		err := i.Initialize(ctx)
		if err != ErrNameNeeded {
			t.Errorf("err should ne ErrNameNeeded but got %v", err)
		}
	})

	t.Run("valid args", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}
		ctx := context.Background()
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "app")
		ctx = context.WithValue(ctx, "args", []string{"new", "app"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(ctx)
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}

		v := filepath.Join(root, "app")
		if _, err := os.Stat(v); os.IsNotExist(err) {
			t.Errorf("should have created %v", v)
		}
	})

	t.Run("force disabled", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		err = os.MkdirAll(filepath.Join(root, "app"), 0777)
		if err != nil {
			t.Fatal("could not create dir")
		}
		ctx := context.Background()
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "app")
		ctx = context.WithValue(ctx, "args", []string{"new", "app"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(ctx)
		if err != ErrFolderExists {
			t.Errorf("err should be ErrFolderExists but got %v", err)
		}
	})

	t.Run("force enabled", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		err = os.MkdirAll(filepath.Join(root, "app"), 0777)
		if err != nil {
			t.Fatal("could not create dir")
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "root", root)
		ctx = context.WithValue(ctx, "name", "app")
		ctx = context.WithValue(ctx, "args", []string{"new", "app"})
		ctx = context.WithValue(ctx, "folder", filepath.Join(root, "app"))

		i := Initializer{force: true}
		err = i.Initialize(ctx)
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}
	})

}
