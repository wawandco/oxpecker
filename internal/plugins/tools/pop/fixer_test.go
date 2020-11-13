package pop

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFix(t *testing.T) {
	t.Run("DataBaseExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}
		content := []byte("")
		err = ioutil.WriteFile("database.yml", content, 0755)
		if err != nil {
			t.Fatal("could not create file")
		}
		err = os.MkdirAll("config/", 0755)

		f := Fixer{}
		err = f.Fix()
		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}
	})
	t.Run("DataBaseNotExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		f := Fixer{}
		err = f.Fix()
		if err != ErrDatabaseNotExist {
			t.Fatalf("error should be %v, got %v", ErrDatabaseNotExist, err)
		}
	})
}

func TestFileExists(t *testing.T) {
	t.Run("DataBaseExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}
		content := []byte("")
		err = ioutil.WriteFile("database.yml", content, 0755)
		if err != nil {
			t.Fatal("could not create file")
		}

		f := Fixer{}
		_, err = f.fileExists(".")
		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}
	})
	t.Run("DataBaseNotExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		f := Fixer{}
		_, err = f.fileExists(".")
		if err != ErrDatabaseNotExist {
			t.Fatalf("error should be %v, got %v", ErrDatabaseNotExist, err)
		}
	})
}

func TestMoveFile(t *testing.T) {
	t.Run("DataBaseExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}
		content := []byte("")
		err = ioutil.WriteFile("database.yml", content, 0755)
		if err != nil {
			t.Fatal("could not create file")
		}
		err = os.MkdirAll("config/", 0755)

		f := Fixer{}
		err = f.moveFile()
		if err != nil {
			t.Fatalf("error should be nill, got %v", err)
		}
	})
	t.Run("DataBaseNotExist", func(t *testing.T) {
		err := os.Chdir(t.TempDir())
		if err != nil {
			t.Fatal("could not move to tmp folder")
		}

		err = os.MkdirAll("config/", 0755)

		f := Fixer{}
		err = f.moveFile()
		if err == nil {
			t.Fatalf("error should be %v, got nil", err)
		}
	})
}
