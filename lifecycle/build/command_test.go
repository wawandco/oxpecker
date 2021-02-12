package build

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	b := &Command{}

	t.Run("Unset", func(t *testing.T) {
		err := b.setenv()
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "production" {
			t.Errorf("GO_ENV should have been production, got %v", env)
		}
	})

	t.Run("Set to development", func(t *testing.T) {
		err := os.Setenv("GO_ENV", "development")
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		err = b.setenv()
		if err != nil {
			t.Errorf("err should be nil, got %v", err)
		}

		env := os.Getenv("GO_ENV")
		if env != "development" {
			t.Errorf("GO_ENV should have been %v, got %v", "development", env)
		}

	})
}
