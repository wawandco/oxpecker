package standard

import "testing"

func Test_ParseFlags_Empty(t *testing.T) {
	c := &Plugin{}
	c.ParseFlags([]string{})

	if c.output != "" {
		t.Errorf("output should be empty, was `%s`", c.output)
	}
}

func Test_ParseFlags_Value(t *testing.T) {
	c := &Plugin{}
	c.ParseFlags([]string{"-o", "something"})

	expected := "something"
	if c.output != expected {
		t.Errorf("output should be `%s`, was `%s`", expected, c.output)
	}
}
