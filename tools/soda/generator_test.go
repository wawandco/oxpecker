package soda

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Generate(t *testing.T) {
	g := Generator{}

	t.Run("generate fizz migration", func(t *testing.T) {
		dir := t.TempDir()

		if err := g.Generate(context.Background(), dir, []string{"generate", "migration", "users", "--type=fizz"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		match, err := filepath.Glob(filepath.Join(dir, "migrations", "*users.*.fizz"))
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

	t.Run("generate sql migration", func(t *testing.T) {
		dir := t.TempDir()

		if err := g.Generate(context.Background(), dir, []string{"generate", "migration", "company", "--type=sql"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		match, err := filepath.Glob(filepath.Join(dir, "migrations", "*companies.*.sql"))
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

	t.Run("generate invalid migration should error", func(t *testing.T) {
		if err := g.Generate(context.Background(), t.TempDir(), []string{"generate", "migration", "company", "--type=invalid"}); err == nil {
			t.Errorf("should be error, but got %v", err)
		}
	})

	t.Run("generate migration without type should generate fizz", func(t *testing.T) {
		dir := t.TempDir()

		if err := g.Generate(context.Background(), dir, []string{"generate", "migration", "templates"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		match, err := filepath.Glob(filepath.Join(dir, "migrations", "*templates.*.fizz"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_templates.down.fizz") {
			t.Error("'templates.up.fizz' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_templates.up.fizz") {
			t.Error("'templates.down.fizz' file does not exists on the path")
		}
	})

	t.Run("generate migration with empty type should generate fizz", func(t *testing.T) {
		dir := t.TempDir()

		if err := g.Generate(context.Background(), dir, []string{"generate", "migration", "invoices", "--type"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		match, err := filepath.Glob(filepath.Join(dir, "migrations", "*invoices.*.fizz"))
		if err != nil {
			t.Errorf("searching for files should not error, but got %v", err)
		}

		if len(match) == 0 {
			t.Error("migration files does not exists on the path")
		}

		if !strings.HasSuffix(match[0], "_invoices.down.fizz") {
			t.Error("'invoices.up.fizz' file does not exists on the path")
		}

		if !strings.HasSuffix(match[1], "_invoices.up.fizz") {
			t.Error("'invoices.down.fizz' file does not exists on the path")
		}
	})

	t.Run("generate fizz migration with args", func(t *testing.T) {
		dir := t.TempDir()

		if err := g.Generate(context.Background(), dir, []string{"generate", "migration", "users", "description:string", "quantity:int"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		match, err := filepath.Glob(filepath.Join(dir, "migrations", "*users.*.fizz"))
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

		// Validating existence of the attributes
		dropFile, err := ioutil.ReadFile(match[0])
		if err != nil {
			t.Error("reading migration down file error")
		}

		if string(dropFile) != `drop_table("users")` {
			t.Error(`unexpected content, file should contain 'drop_table("users")'`)
		}

		createFile, err := ioutil.ReadFile(match[1])
		if err != nil {
			t.Error("reading migration down file error")
		}

		upData := string(createFile)
		shouldContain := []string{`create_table("users")`, "t.Column", "description", "quantity", "string", "integer"}
		for _, contain := range shouldContain {
			if !strings.Contains(upData, contain) {
				t.Errorf("unexpected content, file should contain '%s'", contain)
			}
		}
	})
}
