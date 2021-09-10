package envy_test

import (
	"context"
	"os"
	"testing"

	"github.com/gobuffalo/envy"
	envypl "github.com/wawandco/ox/tools/envy"
)

func TestDeveloper(t *testing.T) {
	t.Run("GO_ENV Not set", func(t *testing.T) {
		d := envypl.Developer{}
		err := d.BeforeDevelop(context.Background(), "", []string{})
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "development" {
			t.Errorf("GO_ENV should be 'development', got %v", env)
		}

		env = envy.Get("GO_ENV", "")
		if env != "development" {
			t.Errorf("GO_ENV should be 'development', got %v", env)
		}
	})

	t.Run("GO_ENV Set previously", func(t *testing.T) {
		os.Setenv("GO_ENV", "somethingelse")
		d := envypl.Developer{}
		err := d.BeforeDevelop(context.Background(), "", []string{})
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "somethingelse" {
			t.Errorf("GO_ENV should be 'somethingelse', got %v", env)
		}

		env = envy.Get("GO_ENV", "")
		if env != "somethingelse" {
			t.Errorf("GO_ENV should be 'somethingelse', got %v", env)
		}
	})
}
