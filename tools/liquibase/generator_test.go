package liquibase

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratorRun(t *testing.T) {
	t.Run("incomplete arguments", func(t *testing.T) {
		g := Generator{}
		err := g.Generate(context.Background(), "", []string{"a", "b"})
		if err != ErrNameArgMissing {
			t.Errorf("err should be %v, got %v", ErrNameArgMissing, err)
		}
	})

	t.Run("simple", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		g := Generator{mockTimestamp: "12345"}
		err = g.Generate(context.Background(), root, []string{"generate", "migration", "aaa"})
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "12345-aaa.xml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Error("should have created the file in the root")
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if content := string(d); !strings.Contains(content, "12345-aaa") {
			t.Errorf("file content %v should contain %v", content, "12345-aaa")
		}
	})

	t.Run("folder", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		g := Generator{mockTimestamp: "12345"}
		err = g.Generate(context.Background(), root, []string{"generate", "migration", "folder/is/here/aaa"})
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "folder", "is", "here", "12345-aaa.xml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Error("should have created the file in the root")
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if content := string(d); !strings.Contains(content, "12345-aaa") {
			t.Errorf("file content %v should contain %v", content, "12345-aaa")
		}
	})

	t.Run("folder exists", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join("folder", "is", "here"), 0755)
		if err != nil {
			t.Fatal("could not create the folder")
		}

		g := Generator{mockTimestamp: "12345"}
		err = g.Generate(context.Background(), root, []string{"generate", "migration", "folder/is/here/aaa"})
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "folder", "is", "here", "12345-aaa.xml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Error("should have created the file in the root")
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if content := string(d); !strings.Contains(content, "12345-aaa") {
			t.Errorf("file content %v should contain %v", content, "12345-aaa")
		}
	})

	t.Run("different base", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		err = os.MkdirAll(filepath.Join("folder", "is", "here"), 0755)
		if err != nil {
			t.Fatal("could not create the folder")
		}

		g := Generator{
			mockTimestamp: "12345",
			baseFolder:    "migrations",
		}

		err = g.Generate(context.Background(), root, []string{"generate", "migration", "aaa"})
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		path := filepath.Join(root, "migrations", "12345-aaa.xml")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			t.Error("should have created the file in the root/migrations folder")
		}

		d, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		if content := string(d); !strings.Contains(content, "12345-aaa") {
			t.Errorf("file content %v should contain %v", content, "12345-aaa")
		}
	})
}

func TestGeneratorComposeName(t *testing.T) {
	t.Run("Valid name", func(t *testing.T) {
		g := Generator{}

		filename, err := g.composeFilename("addDevices", "composename")
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		expected := "composename-add_devices.xml"
		if filename != expected {
			t.Errorf("filename should be %v, got %v", expected, filename)
		}
	})

	t.Run("Invalid name", func(t *testing.T) {
		g := Generator{}
		_, err := g.composeFilename(".", "composename")
		if err != ErrInvalidName {
			t.Errorf("err should be ErrInvalidName, got %v", err)
		}

		_, err = g.composeFilename("/", "composename")
		if err != ErrInvalidName {
			t.Errorf("err should be ErrInvalidName, got %v", err)
		}
	})
}
