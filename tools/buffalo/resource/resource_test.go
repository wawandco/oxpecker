package resource

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_GenerateActions(t *testing.T) {
	dir := t.TempDir()
	args := []string{"companies"}

	if err := os.MkdirAll(filepath.Join(dir, "app", "actions"), os.ModePerm); err != nil {
		t.Errorf("creating actions directory should not error, but got %v", err)
	}

	resource := New(dir, args)
	if err := resource.GenerateActions(); err != nil {
		t.Errorf("should not be error, but got %v", err)
	}

	// Checking Action Files
	if _, err := os.Stat(filepath.Join(dir, "app", "actions", "companies.go")); os.IsNotExist(err) {
		t.Error("'companies.go' file does not exists on the path")
	}

	if _, err := os.Stat(filepath.Join(dir, "app", "actions", "companies_test.go")); os.IsNotExist(err) {
		t.Error("'companies_test.go' file does not exists on the path")
	}

	// Validating existence of the attributes
	companyDir := filepath.Join(dir, "app", "actions", "companies.go")
	data, err := ioutil.ReadFile(companyDir)
	if err != nil {
		t.Error("reading file error")
	}

	stringData := string(data)
	shouldContain := []string{
		"func (v CompaniesResource) List(c buffalo.Context) error {",
		"func (v CompaniesResource) New(c buffalo.Context) error {",
		"func (v CompaniesResource) Show(c buffalo.Context) error {",
		"func (v CompaniesResource) Update(c buffalo.Context) error {",
		"func (v CompaniesResource) Destroy(c buffalo.Context) error {",
		"func (v CompaniesResource) Create(c buffalo.Context) error {",
	}
	for _, contain := range shouldContain {
		if !strings.Contains(stringData, contain) {
			t.Errorf("unexpected content, file should contain '%s'", contain)
		}
	}
}

func Test_GenerateMigrations(t *testing.T) {
	dir := t.TempDir()
	args := []string{"events"}

	if err := os.MkdirAll(filepath.Join(dir, "migrations"), os.ModePerm); err != nil {
		t.Errorf("creating actions directory should not error, but got %v", err)
	}

	resource := New(dir, args)
	if err := resource.GenerateMigrations(); err != nil {
		t.Errorf("should not be error, but got %v", err)
	}

	// Validating Files existence
	match, err := filepath.Glob(filepath.Join(dir, "migrations", "*events.*.fizz"))
	if err != nil {
		t.Errorf("searching for files should not error, but got %v", err)
	}

	if len(match) == 0 {
		t.Error("migration files does not exists on the path")
	}

	if !strings.HasSuffix(match[0], "_events.down.fizz") {
		t.Error("'events.up.fizz' file does not exists on the path")
	}

	if !strings.HasSuffix(match[1], "_events.up.fizz") {
		t.Error("'events.down.fizz' file does not exists on the path")
	}
}

func Test_GenerateModels(t *testing.T) {
	dir := t.TempDir()
	args := []string{"users", "first_name", "last_name", "email"}
	modelsPath := filepath.Join(dir, "app", "models")
	if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
		t.Errorf("creating templates folder should not be error, but got %v", err)
	}

	resource := New(dir, args)
	if err := resource.GenerateModel(); err != nil {
		t.Errorf("should not be error, but got %v", err)
	}

	// Validating Files existence
	if _, err := os.Stat(filepath.Join(dir, "app", "models", "user.go")); os.IsNotExist(err) {
		t.Error("'users.go' file does not exists on the path")
	}

	if _, err := os.Stat(filepath.Join(dir, "app", "models", "user_test.go")); os.IsNotExist(err) {
		t.Error("'users_test.go' file does not exists on the path")
	}

	// Validating existence of the attributes
	userFilePath := filepath.Join(modelsPath, "user.go")
	data, err := ioutil.ReadFile(userFilePath)
	if err != nil {
		t.Error("reading file error")
	}

	stringData := string(data)
	shouldContain := []string{"ID", "CreatedAt", "UpdatedAt", "FirstName", "LastName", "Email", "time.Time", "string"}
	for _, contain := range shouldContain {
		if !strings.Contains(stringData, contain) {
			t.Errorf("unexpected content, file should contain '%s'", contain)
		}
	}
}

func Test_GenerateTemplates(t *testing.T) {
	dir := t.TempDir()
	args := []string{"cars", "model", "brand"}
	modelsPath := filepath.Join(dir, "app", "templates")
	if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
		t.Errorf("creating templates folder should not be error, but got %v", err)
	}

	resource := New(dir, args)
	if err := resource.GenerateTemplates(); err != nil {
		t.Errorf("should not be error, but got %v", err)
	}

	// Validating Files existence
	templateFolder := filepath.Join(dir, "app", "templates", "cars")
	templates := []string{"index", "new", "edit", "show", "form"}

	for _, tmpl := range templates {
		if _, err := os.Stat(filepath.Join(templateFolder, tmpl+".plush.html")); os.IsNotExist(err) {
			t.Error("'index.plush.html' file does not exists on the path")
		}
	}
}
