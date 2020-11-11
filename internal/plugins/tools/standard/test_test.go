package standard

import (
	"context"
	"os"
	"testing"
)

func Test_RunBeforeTest(t *testing.T) {
	p := &Plugin{}
	ctx := context.Background()
	err := p.RunBeforeTest(ctx, "", []string{})

	if err != nil {
		t.Errorf("Should not error, got %v", err)
	}

	env := os.Getenv("GO_ENV")
	if env != "test" {
		t.Errorf("GO_ENV should is %v should be %v", env, "test")
	}
}
