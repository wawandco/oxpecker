package grift

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func Test_ActionGenerator(t *testing.T) {
	g := Generator{}

	t.Run("generate task", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "tasks")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "task", "simple"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		if !g.exists(filepath.Join(modelsPath, "simple.go")) {
			t.Error("'simple.go' file does not exists on the path")
		}
	})
	t.Run("generate task and checking the content", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "tasks")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "task", "simple"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		if !g.exists(filepath.Join(modelsPath, "simple.go")) {
			t.Error("'simple.go' file does not exists on the path")
		}

		content, err := ioutil.ReadFile(filepath.Join(modelsPath, "simple.go"))
		if err != nil {
			log.Fatal(err)
		}
		text := string(content)
		matched, err := regexp.MatchString(`"simple", func`, text)

		if !matched {
			t.Fatalf("File's content is not correct, %v", err)
		}
	})
}
