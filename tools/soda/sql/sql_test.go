package sql

import (
	"path/filepath"
	"strings"
	"testing"
)

func Test_SQL_Create(t *testing.T) {
	s := Creator{}

	t.Run("generate migration files", func(t *testing.T) {
		dir := t.TempDir()
		name := "users"
		args := []string{}

		if err := s.Create(dir, name, args); err != nil {
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
		name := "company"
		args := []string{}

		if err := s.Create(dir, name, args); err != nil {
			t.Errorf("creating migration files should not be error, but got %v", err)
		}

		// Validating files existence
		match, err := filepath.Glob(filepath.Join(dir, "*company.*.sql"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "company.down.sql") {
			t.Error("'company.up.sql' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "company.up.sql") {
			t.Error("'company.down.sql' file does not exists on the path")
		}
	})
}
