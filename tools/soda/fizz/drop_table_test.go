package fizz

import (
	"strings"
	"testing"
)

func Test_Drop_Table(t *testing.T) {
	ac := dropTable{}
	t.Run("no table name", func(t *testing.T) {
		_, _, err := ac.GenerateFizz("drop_table_", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("with table name", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("drop_table_users", []string{})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `drop_table("users")`
		expectedDown := `create_table("users")`

		if up != expectedUP {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if !strings.Contains(down, expectedDown) {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Drop_Table_Matches(t *testing.T) {
	ac := dropTable{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "drop_table_", expected: true},
		{name: "drop_table", expected: false},
		{name: "drop_table_users", expected: true},
	}

	for _, c := range cases {
		matchs := ac.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}
