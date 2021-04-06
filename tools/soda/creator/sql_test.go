package creator

import (
	"path/filepath"
	"strings"
	"testing"
)

func Test_SQL_Create(t *testing.T) {
	s := SQLCreator{}

	t.Run("generate migration files", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"users"}

		if err := s.Create(dir, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*users.*.sql"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_users.down.sql") {
			t.Error("'users.up.sql' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_users.up.sql") {
			t.Error("'users.down.sql' file does not exists on the path")
		}
	})

	t.Run("generate migration singularized name", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"company"}

		if err := s.Create(dir, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*companies.*.sql"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_companies.down.sql") {
			t.Error("'companies.up.sql' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_companies.up.sql") {
			t.Error("'companies.down.sql' file does not exists on the path")
		}
	})
}
