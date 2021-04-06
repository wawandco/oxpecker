package creator

import (
	"path/filepath"
	"strings"
	"testing"
)

func Test_Fizz_Create(t *testing.T) {
	f := FizzCreator{}

	t.Run("generate migration files", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"users"}

		if err := f.Create(dir, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*users.*.fizz"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_users.down.fizz") {
			t.Error("'users.up.fizz' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_users.up.fizz") {
			t.Error("'users.down.fizz' file does not exists on the path")
		}
	})

	t.Run("generate migration singularized name", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"company"}

		if err := f.Create(dir, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*companies.*.fizz"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_companies.down.fizz") {
			t.Error("'companies.up.fizz' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_companies.up.fizz") {
			t.Error("'companies.down.fizz' file does not exists on the path")
		}
	})
}
