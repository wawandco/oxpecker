package yarn

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

// TestBuildCmd attempts to run yarn install if it finds yarn.lock
func TestBuildCmd(t *testing.T) {
	err := os.Chdir(t.TempDir())
	if err != nil {
		t.Error(err)
	}

	tcases := []struct {
		commandNil  bool
		failMessage string
		beforeFn    func()
	}{
		{
			commandNil:  false,
			failMessage: "command should not be nil",
			beforeFn: func() {
				err := ioutil.WriteFile("yarn.lock", []byte{}, 0600)
				if err != nil {
					t.Error(err)
				}
			},
		},

		{
			commandNil:  true,
			failMessage: "command should be nil",
			beforeFn:    func() {},
		},
	}

	p := &Plugin{}
	for _, tcase := range tcases {
		err := os.Remove("yarn.lock")
		if err != nil && !os.IsNotExist(err) {
			t.Error(err)
			break
		}

		tcase.beforeFn()

		c := p.buildCmd(context.Background())
		cond := tcase.commandNil && c != nil
		cond = cond || !tcase.commandNil && c == nil

		if cond {
			t.Error(tcase.failMessage)
		}
	}
}
