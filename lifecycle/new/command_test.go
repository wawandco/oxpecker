package new_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/oxpecker/lifecycle/new"
	"github.com/wawandco/oxpecker/plugins"
)

func TestRun(t *testing.T) {
	root := t.TempDir()
	err := os.Chdir(root)
	if err != nil {
		t.Error("could not change to temp directory")
	}

	pl := &new.Command{}
	tinit := &Tinit{}
	pl.Receive([]plugins.Plugin{tinit})

	err = pl.Run(context.Background(), root, []string{})
	if err == nil {
		t.Error("should return an error")
	}

	err = pl.Run(context.Background(), root, []string{"app"})
	if err != nil {
		t.Errorf("should not return and error, got: %v", err)
	}

	//Should create the folder
	fi, err := os.Stat(filepath.Join(root, "app"))
	if err != nil {
		t.Errorf("should not return and error, got: %v", err)
	}

	if !fi.IsDir() {
		t.Errorf("should be a folder, got a file")
	}

	if !tinit.called {
		t.Errorf("should have called initializer")
	}

	if !tinit.afterCalled {
		t.Errorf("should have called afterinitialize")
	}

	if tinit.root != filepath.Join(root, "app") {
		t.Errorf("should call initializer with root being: %v", filepath.Join(root, "app"))
	}
}

func TestFolderName(t *testing.T) {
	tcases := []struct {
		args     []string
		expected string
	}{
		{[]string{"aaa"}, "aaa"},
		{[]string{"something/aaa"}, "aaa"},
		{[]string{"something\\aaa"}, "something\\aaa"},
	}

	pl := &new.Command{}
	for _, tcase := range tcases {
		name := pl.FolderName(tcase.args)
		if name != tcase.expected {
			t.Errorf("should return %v got %v", tcase.expected, name)
		}
	}

}
