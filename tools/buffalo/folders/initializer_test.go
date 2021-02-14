package folders

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestInitializerRun(t *testing.T) {

	t.Run("valid args", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Fatal("could not move to temp dir")
		}

		var dx sync.Map
		dx.Store("root", root)
		dx.Store("name", "app")
		dx.Store("args", []string{"new", "app"})
		dx.Store("folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(context.Background(), &dx)
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
		var dx sync.Map

		err := i.Initialize(context.Background(), &dx)
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

		var dx sync.Map
		dx.Store("root", root)
		dx.Store("name", "app")
		dx.Store("args", []string{"new", "app"})
		dx.Store("folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(context.Background(), &dx)
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

		var dx sync.Map
		dx.Store("root", root)
		dx.Store("name", "app")
		dx.Store("args", []string{"new", "app"})
		dx.Store("folder", filepath.Join(root, "app"))

		i := Initializer{}
		err = i.Initialize(context.Background(), &dx)
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

		var dx sync.Map
		dx.Store("root", root)
		dx.Store("name", "app")
		dx.Store("args", []string{"new", "app"})
		dx.Store("folder", filepath.Join(root, "app"))

		i := Initializer{force: true}
		err = i.Initialize(context.Background(), &dx)
		if err != nil {
			t.Errorf("err should be nil but got %v", err)
		}
	})

}
