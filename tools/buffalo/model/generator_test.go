package model

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Generate(t *testing.T) {
	g := Generator{}

	t.Run("generate model", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "models")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "model", "users"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		if !g.exists(filepath.Join(modelsPath, "user.go")) {
			t.Error("'user.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(modelsPath, "user_test.go")) {
			t.Error("'user_test.go' file does not exists on the path")
		}
	})

	t.Run("generate model validating model name with underscore", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "models")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "model", "organizational_model"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		userFilePath := filepath.Join(modelsPath, "organizational_model.go")
		data, err := ioutil.ReadFile(userFilePath)
		if err != nil {
			t.Error("reading file error")
		}

		stringData := string(data)
		shouldContain := []string{"OrganizationalModel", "OrganizationalModels", "[]OrganizationalModel"}
		for _, contain := range shouldContain {
			if !strings.Contains(stringData, contain) {
				t.Errorf("unexpected content, file should contain '%s'", contain)
			}
		}
	})

	t.Run("generate model with args", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "models")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "model", "users", "description:string", "quantity:int"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		userFilePath := filepath.Join(modelsPath, "user.go")
		if !g.exists(userFilePath) {
			t.Error("'user.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(modelsPath, "user_test.go")) {
			t.Error("'user_test.go' file does not exists on the path")
		}

		// Validating existence of the attributes
		data, err := ioutil.ReadFile(userFilePath)
		if err != nil {
			t.Error("reading file error")
		}

		stringData := string(data)
		shouldContain := []string{"ID", "CreatedAt", "UpdatedAt", "Description", "Quantity", "uuid.UUID", "time.Time", "string", "int"}
		for _, contain := range shouldContain {
			if !strings.Contains(stringData, contain) {
				t.Errorf("unexpected content, file should contain '%s'", contain)
			}
		}
	})
}
