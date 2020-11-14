package packr

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestRunBeforeBuild(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		dir := os.TempDir()
		err := os.Chdir(dir)
		if err != nil {
			t.Error(err)
		}

		err = ioutil.WriteFile("go.mod", []byte("module sample/thing"), 0644)
		if err != nil {
			t.Error(err)
		}

		p := &Plugin{}
		err = p.RunBeforeBuild(context.Background(), dir, []string{})
		if err != nil {
			t.Error(err)
		}

		_, err = os.Stat("thing.go")
		if os.IsNotExist(err) {
			t.Fatal("thing.go should be generated!")
		}

		content, err := ioutil.ReadFile("thing.go")
		if err != nil {
			t.Error(err)
		}

		if !bytes.Contains(content, []byte("package thing")) {
			t.Errorf("file should contain `package thing`")
		}
	})

	t.Run("Exists", func(t *testing.T) {
		dir := os.TempDir()
		err := os.Chdir(dir)
		if err != nil {
			t.Error(err)
		}

		err = ioutil.WriteFile("go.mod", []byte("module sample/thing"), 0644)
		if err != nil {
			t.Error(err)
		}

		err = ioutil.WriteFile("thing.go", []byte("package other"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		p := &Plugin{}
		err = p.RunBeforeBuild(context.Background(), dir, []string{})
		if err != nil {
			t.Error(err)
		}

		content, err := ioutil.ReadFile("thing.go")
		if err != nil {
			t.Error(err)
		}

		if !bytes.Contains(content, []byte("package other")) {
			t.Errorf("file should contain `package other`")
		}
	})
}
