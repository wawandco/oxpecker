//go:build integration
// +build integration

package cli_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/cli"
)

func TestNewApp(t *testing.T) {
	dir := t.TempDir()
	os.Chdir(dir)

	err := cli.Run(context.Background(), []string{"ox", "new", "coke"})
	if err != nil {
		t.Fatalf("error running new command: %v", err)
	}

	files := [][]string{
		{dir, "coke"},
		{dir, "coke", "go.mod"},
		{dir, "coke", "embed.go"},
	}

	for _, f := range files {
		file := filepath.Join(f...)
		if _, err := os.Stat(file); err == nil {
			continue
		}

		t.Fatalf("did not find: %v", file)
	}

}
