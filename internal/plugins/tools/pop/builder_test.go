package pop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_findConfig_Exists_config_folder(t *testing.T) {
	dir := t.TempDir()
	err := os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}

	config := `
development:
	pool: 2
	url: something
	`
	err = os.Mkdir("config", 0700)
	if err != nil {
		t.Errorf("should be no error creating folder, got %v", err)
	}
	fmt.Println(dir)

	err = ioutil.WriteFile(filepath.Join("config", "database.yml"), []byte(config), 0600)
	if err != nil {
		t.Error(err)
	}

	p := &Plugin{}
	b, err := p.findConfig()
	if err != nil {
		t.Errorf("should be no error but got %v", err)
	}

	if !bytes.Contains(b, []byte("url: something")) {
		t.Error("should contain url:something")
	}
}

func Test_findConfig_notThere(t *testing.T) {
	dir := t.TempDir()
	err := os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}

	p := &Plugin{}
	_, err = p.findConfig()
	if err == nil {
		t.Errorf("should be error but got nil")
	}
}

func Test_RunAfterBuild(t *testing.T) {
	dir := t.TempDir()
	err := os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}

	err = os.MkdirAll("config", 0700)
	if err != nil {
		t.Error(err)
	}

	path := filepath.Join("config", "gen_database.go")
	err = ioutil.WriteFile(path, []byte{}, 0600)
	if err != nil {
		t.Error(err)
	}

	p := &Plugin{}
	err = p.RunAfterBuild(dir, []string{})
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(path); err == nil {
		t.Errorf("%v should not exist after build", path)
	}
}

func Test_RunAfterBuild_NoFile(t *testing.T) {
	dir := t.TempDir()
	err := os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}

	p := &Plugin{}
	err = p.RunAfterBuild(dir, []string{})
	if err != nil {
		t.Error(err)
	}

	path := filepath.Join("config", "gen_database.go")
	if _, err := os.Stat(path); err == nil {
		t.Errorf("%v should not exist after build", path)
	}
}
