package soda

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wawandco/ox/plugins"
	"github.com/wawandco/ox/tools/soda/fizz"
	"github.com/wawandco/ox/tools/soda/sql"
)

func Test_Generate(t *testing.T) {
	g := Generator{}
	g.Receive([]plugins.Plugin{&fizz.Creator{}, &sql.Creator{}})

	t.Run("generate fizz migration", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"generate", "migration", "users", "--type=fizz"}

		g.ParseFlags(args)
		if err := g.Generate(context.Background(), dir, args); err != nil {
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
		args := []string{"generate", "migration", "company", "--type=sql"}
		g.ParseFlags(args)

		if err := g.Generate(context.Background(), dir, args); err != nil {
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
		args := []string{"generate", "migration", "company", "--type=invalid"}
		g.ParseFlags(args)

		if err := g.Generate(context.Background(), t.TempDir(), args); err == nil {
			t.Errorf("should be error, but got %v", err)
		}
	})

	t.Run("generate migration without type should generate fizz", func(t *testing.T) {
		dir := t.TempDir()
		args := []string{"generate", "migration", "templates"}
		g.ParseFlags(args)

		if err := g.Generate(context.Background(), dir, args); err != nil {
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
		args := []string{"generate", "migration", "invoices", "--type"}
		g.ParseFlags(args)

		if err := g.Generate(context.Background(), dir, args); err != nil {
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
		args := []string{"generate", "migration", "users", "description:string", "quantity:int"}
		g.ParseFlags(args)

		if err := g.Generate(context.Background(), dir, args); err != nil {
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
