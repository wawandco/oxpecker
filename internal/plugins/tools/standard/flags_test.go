package standard

import "testing"

func Test_ParseFlags_Empty(t *testing.T) {
	c := &Plugin{}
	err := c.ParseFlags([]string{})
	if err != nil {
		t.Errorf("should not err on parseFlags empty")
	}

	if c.output != "" {
		t.Errorf("output should be empty, was `%s`", c.output)
	}
}

func Test_ParseFlags_Value(t *testing.T) {
	c := &Plugin{}
	err := c.ParseFlags([]string{"-o", "something"})
	if err != nil {
		t.Errorf("should not err on parseFlags with value")
	}

	expected := "something"
	if c.output != expected {
		t.Errorf("output should be `%s`, was `%s`", expected, c.output)
	}
}
