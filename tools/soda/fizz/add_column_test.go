package fizz

import (
	"strings"
	"testing"
)

func Test_Add_Column(t *testing.T) {
	ac := addColumn{}
	t.Run("no column found", func(t *testing.T) {
		_, _, err := ac.GenerateFizz("add_description_to_templates", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("1 argument", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("adding_money_to_bank", []string{"money:int"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `add_column("bank", "money", "integer", {})`
		expectedDown := `drop_column("bank", "money")`

		if up != expectedUP {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})

	t.Run("multiple arguments", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("adding_columns_to_table", []string{"money:int", "name", "email"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		upContains := []string{`add_column("table", "name", "string", {})`, `add_column("table", "email", "string", {})`, `add_column("table", "money", "integer", {})`}
		downContains := []string{`drop_column("table", "money")`, `drop_column("table", "name")`, `drop_column("table", "email")`}

		for _, v := range upContains {
			if !strings.Contains(up, v) {
				t.Errorf("expected %v but got %v", v, up)
			}
		}

		for _, v := range downContains {
			if !strings.Contains(down, v) {
				t.Errorf("expected %v but got %v", v, down)
			}
		}
	})
}

func Test_Add_Column_Matches(t *testing.T) {
	ac := addColumn{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "add_description_to_templates", expected: true},
		{name: "adding_money_to_bank", expected: true},
		{name: "adding_users", expected: false},
		{name: "add_companies", expected: false},
	}

	for _, c := range cases {
		matchs := ac.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}
