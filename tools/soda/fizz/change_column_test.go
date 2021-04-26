package fizz

import "testing"

func Test_Change_Column(t *testing.T) {
	cc := changeColumn{}

	t.Run("no new column arg", func(t *testing.T) {
		_, _, err := cc.GenerateFizz("change_companies_name", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("with new column arg", func(t *testing.T) {
		up, down, err := cc.GenerateFizz("change_companies_price", []string{"float"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `change_column("companies", "price", "decimal", {})`
		expectedDown := `change_column("companies", "price", "string", {})`

		if up != expectedUP {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})

	t.Run("with new column null arg", func(t *testing.T) {
		up, down, err := cc.GenerateFizz("change_companies_price", []string{"nulls.Int"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `change_column("companies", "price", "integer", {null: true})`
		expectedDown := `change_column("companies", "price", "string", {})`

		if up != expectedUP {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Change_Column_Matches(t *testing.T) {
	cc := changeColumn{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "change", expected: false},
		{name: "change_users", expected: false},
		{name: "change_companies_name", expected: true},
	}

	for _, c := range cases {
		matchs := cc.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}
