package action

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
			Module: "oosss/myapp",
			Folder: filepath.Join(root, "myapp"),
		}

		err = i.Initialize(ctx, options)
		if err != nil {
			t.Fatalf("error should be nil, got %v", err)
		}

		tcases := []struct {
			path    string
			content []string
		}{
			{
				path: filepath.Join(root, "myapp", "app", "actions", "actions.go"),
				content: []string{
					`r = render.Engine`,
					`"oosss/myapp/app/render"`,
				},
			},

			{
				path: filepath.Join(root, "myapp", "app", "actions", "actions_test.go"),
				content: []string{
					`package actions_test`,
					`"oosss/myapp/app"`,
					`func Test_ActionSuite(t *testing.T) {`,
				},
			},
		}

		for _, tcase := range tcases {
			bm, err := ioutil.ReadFile(tcase.path)
			if err != nil {
				t.Fatal("should have created the file:", tcase.path)
			}

			for _, cnt := range tcase.content {
				if !bytes.Contains(bm, []byte(cnt)) {
					t.Errorf("%v does not contain `%v`", tcase.path, cnt)
				}
			}
		}

	})
}
