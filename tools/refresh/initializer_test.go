package refresh

import (
	"context"
	"os"
	"testing"
)

func TestInitializer(t *testing.T) {
	t.Run("BuffaloFileDoesNotExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})

		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		_, err = os.Stat(root)

		if os.IsNotExist(err) {
			t.Fatalf("Did not create file ")
		}

	})
	t.Run("BuffaloFileExist", func(t *testing.T) {

		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		rootYml := root + "/.buffalo.dev.yml"
		_, err = os.Create(rootYml)
		if err != nil {
			t.Fatalf("Problem creating file, %v", err)
		}

		i := Initializer{}

		err = i.Initialize(context.Background(), root, []string{})

		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

	})
}
