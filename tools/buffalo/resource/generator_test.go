package resource

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func Test_Generate(t *testing.T) {
	g := Generator{}
	dir := t.TempDir()

	// Creating folders before hand
	folders := []string{"app/actions", "app/models", "app/templates", "migrations"}
	for _, f := range folders {
		actionsPath := filepath.Join(dir, f)
		if err := os.MkdirAll(actionsPath, os.ModePerm); err != nil {
			t.Errorf("creating %s folder should not be error, but got %v", f, err)
		}
	}

	if err := g.Generate(context.Background(), dir, []string{"generate", "resource", "animals", "age:int", "breed", "nationality"}); err != nil {
		t.Errorf("should not be error, but got %v", err)
	}

	// Validating Files existence
	files := []string{
		"app/actions/animals.go",
		"app/actions/animals_test.go",
		"app/models/animal.go",
		"app/models/animal_test.go",
		"app/templates/animals/index.plush.html",
		"app/templates/animals/new.plush.html",
		"app/templates/animals/edit.plush.html",
		"app/templates/animals/show.plush.html",
		"app/templates/animals/form.plush.html",
	}

	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("'%s' file does not exists on the path", path)
		}
	}
}
