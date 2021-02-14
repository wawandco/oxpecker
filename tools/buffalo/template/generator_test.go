package template

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_GenerateTemplate(t *testing.T) {
	g := Generator{}

	t.Run("generate template", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(dir, "app", "templates"), os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.generateTemplate(dir, "amazon"); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}
	})

	t.Run("generate template with subfolder", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(dir, "app", "templates"), os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.generateTemplate(dir, "partials/sidebar"); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}
	})

	t.Run("generate template with subfolder 2", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(dir, "app", "templates"), os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.generateTemplate(dir, "templates/index"); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}
	})

	t.Run("generated template already exists", func(t *testing.T) {
		dir := t.TempDir()
		templatesPath := filepath.Join(dir, "app", "templates")
		if err := os.MkdirAll(templatesPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if _, err := os.Create(filepath.Join(templatesPath, "ebay.plush.html")); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		if err := g.generateTemplate(dir, "ebay"); err == nil {
			t.Errorf("should error, but got %v", err)
		}
	})

	t.Run("generate template when templates folder do not exists", func(t *testing.T) {
		if err := g.generateTemplate("", "user"); err == nil {
			t.Errorf("should error, but got %v", err)
		}
	})
}

func Test_GenerateFilePath(t *testing.T) {
	g := Generator{}
	cases := []struct {
		expected string
		fileName string
		testName string
	}{
		{
			testName: "normal filename",
			fileName: "user",
			expected: "user.plush.html",
		},
		{
			testName: "filename with extension 1",
			fileName: "amazon.plush.html",
			expected: "amazon.plush.html",
		},
		{
			testName: "filename with extension 2",
			fileName: "ebay.html",
			expected: "ebay.plush.html",
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			filePath := g.generateFilePath("", c.fileName)

			if filePath != c.expected {
				t.Errorf("test was expecting %s, got %s", c.expected, filePath)
			}
		})
	}
}
