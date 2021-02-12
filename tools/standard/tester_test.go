package standard

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestRunBeforeTest(t *testing.T) {
	p := &Tester{}
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

func TestCommandArgs(t *testing.T) {
	p := &Tester{}

	tcases := []struct {
		args     []string
		expected string
		name     string
	}{
		{
			args:     []string{"-p", "3"},
			expected: "test -p 3",
			name:     "P present",
		},

		{
			args:     []string{},
			expected: "test -p 1 ./...",
			name:     "blank",
		},

		{
			args:     []string{"./app/actions/..."},
			expected: "test -p 1 ./app/actions/...",
			name:     "only route",
		},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			targs := p.testArgs(tcase.args)
			actual := strings.Join(targs, " ")
			if actual != tcase.expected {
				t.Errorf("should have gotten %v and got %v", tcase.expected, actual)
			}
		})
	}
}
