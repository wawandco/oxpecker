package compiler

import "testing"

func Test_binaryOutput(t *testing.T) {
	c := &Compiler{}
	output := c.binaryOutput("aaa")

	if output != "bin/aaa" {
		t.Errorf("binaryOutput should be %v not %v", "bin/aaa", output)
	}
}
