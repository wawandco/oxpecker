package new_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/lifecycle/new"
	"github.com/wawandco/ox/plugins"
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

	err = pl.Run(context.Background(), root, []string{"new", "app"})
	if err != nil {
		t.Errorf("should not return and error, got: %v", err)
	}

	//Should create the folder
	if !tinit.called {
		t.Errorf("should have called initializer")
	}

	if !tinit.afterCalled {
		t.Errorf("should have called afterinitialize")
	}

	if tinit.root != root {
		t.Errorf("should call initializer with root being: %v", root)
	}

	if tinit.name != "app" {
		t.Errorf("should call initializer with folder being: %v", "app")
	}
	exp := filepath.Join(root, "app")
	if tinit.folder != exp {
		t.Errorf("should call initializer with folder being: %v", exp)
	}
}

func TestFolderName(t *testing.T) {
	tcases := []struct {
		args     []string
		expected string
	}{
		{[]string{"new", "aaa"}, "aaa"},
		{[]string{"new", "something/aaa"}, "aaa"},
		{[]string{"new", "something\\aaa"}, "something\\aaa"},
	}

	pl := &new.Command{}
	for _, tcase := range tcases {
		name := pl.AppName(tcase.args)
		if name != tcase.expected {
			t.Errorf("should return %v got %v", tcase.expected, name)
		}
	}

}
