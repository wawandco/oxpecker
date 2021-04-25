package fizz

import (
	"testing"
)

func Test_Drop_Index(t *testing.T) {
	ac := dropIndex{}

	t.Run("normal drop", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("drop_index_idx_name_from_campaigns", []string{})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `drop_index("campaigns", "idx_name")`
		expectedDown := `add_index("campaigns", "idx_name", {})`

		if up != expectedUP {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Drop_Index_Matches(t *testing.T) {
	ac := dropIndex{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "drop_index", expected: false},
		{name: "drop_index_ix", expected: false},
		{name: "drop_index_idx_name_from_campaigns", expected: true},
	}

	for _, c := range cases {
		matchs := ac.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}
