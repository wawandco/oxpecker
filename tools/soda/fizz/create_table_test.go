package fizz

import (
	"strings"
	"testing"
)

func Test_Create_Table(t *testing.T) {
	ac := createTable{}
	t.Run("no table name", func(t *testing.T) {
		_, _, err := ac.GenerateFizz("create_table_", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("with table name and no args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_table_users", []string{})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `create_table("users")`
		expectedDown := `drop_table("users")`

		if !strings.Contains(up, expectedUP) {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})

	t.Run("with table name and args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_table_users", []string{"email"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP1 := `create_table("users")`
		expectedUP2 := `t.Column("email", "string", {})`
		expectedDown := `drop_table("users")`

		if !strings.Contains(up, expectedUP1) {
			t.Errorf("expected %v but got %v", expectedUP1, up)
		}

		if !strings.Contains(up, expectedUP2) {
			t.Errorf("expected %v but got %v", expectedUP2, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Create_Table_Matches(t *testing.T) {
	ac := createTable{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "create_table_", expected: true},
		{name: "create_table", expected: false},
		{name: "create_table_users", expected: true},
	}

	for _, c := range cases {
		matchs := ac.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}
