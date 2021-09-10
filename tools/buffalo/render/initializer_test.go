package render

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/ox/lifecycle/new"
)

func TestInitializer(t *testing.T) {
	t.Run("CompleteArgs", func(t *testing.T) {
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
		options := new.Options{
			Name:   "myapp",
			Module: "oosss/myapp",
			Folder: filepath.Join(root, "myapp"),
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "myapp", "app", "render", "render.go")
		bm, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal("should have created the file")
		}

		content := []string{
			`package render`,
			`var Engine = render.New(render.Options{`,
			`var Helpers = map[string]interface{}{`,
			`"partialFeeder": base.Templates.FindString,`,
		}

		for _, c := range content {
			if !bytes.Contains(bm, []byte(c)) {
				t.Errorf("`%v` does not contain `%v`", path, c)
			}
		}
	})

}
