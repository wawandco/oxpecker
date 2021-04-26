package fizz

import (
	"path/filepath"
	"strings"
	"testing"
)

func Test_Fizz_Create(t *testing.T) {
	f := Creator{}

	t.Run("generate migration files", func(t *testing.T) {
		dir := t.TempDir()
		name := "users"
		args := []string{}

		if err := f.Create(dir, name, args); err != nil {
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
		name := "company"
		args := []string{}

		if err := f.Create(dir, name, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*company.*.fizz"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_company.down.fizz") {
			t.Error("'companies.up.fizz' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_company.up.fizz") {
			t.Error("'company.down.fizz' file does not exists on the path")
		}
	})
}
