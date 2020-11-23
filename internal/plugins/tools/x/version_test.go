package x

import (
	"testing"

	"github.com/paganotoni/oxpecker/internal/plugins"
)

type testVersioner struct {
	tool string
}

func (tv testVersioner) Name() string {
	return "tool"
}

func (tv testVersioner) ToolName() string {
	return tv.tool
}

func (tv testVersioner) Version() (string, error) {
	return "1.0", nil
}

func TestReceive(t *testing.T) {
	vr := VersionCommand{}
	vr.Receive([]plugins.Plugin{})

	if vr.versioner != nil {
		t.Fatal("versioner should be nil")
	}

	vr.Receive([]plugins.Plugin{
		testVersioner{tool: "ox"},
	})

	if vr.versioner == nil {
		t.Fatal("versioner should not nil")
	}

	vr.versioner = nil
	vr.Receive([]plugins.Plugin{
		testVersioner{tool: "aa"},
	})

	if vr.versioner != nil {
		t.Fatal("versioner should be nil")
	}
}
