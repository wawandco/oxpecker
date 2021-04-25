package fizz

import (
	"errors"
	"testing"
)

func Test_Rename(t *testing.T) {
	r := rename{}
	_, _, err := r.GenerateFizz("rename_something_invalid", []string{})
	if err == nil {
		t.Error("should error but got nil")
	}
}

func Test_Rename_Table(t *testing.T) {
	r := rename{}

	t.Run("invalid table renamer", func(t *testing.T) {
		_, _, err := r.GenerateFizz("rename_table_users", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("normal table renamer", func(t *testing.T) {
		up, down, err := r.GenerateFizz("rename_table_users_to_employees", []string{})
		if err != nil {
			t.Error("should be nil but got error")
		}

		expectedUp := `rename_table("users", "employees")`
		expectedDown := `rename_table("employees", "users")`

		if up != expectedUp {
			t.Errorf("expected %v but got %v", expectedUp, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Rename_Column(t *testing.T) {
	r := rename{}

	t.Run("invalid column renamer", func(t *testing.T) {
		_, _, err := r.GenerateFizz("rename_column_description_from_templates", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("normal column renamer", func(t *testing.T) {
		up, down, err := r.GenerateFizz("rename_column_landing_to_html_from_templates", []string{})
		if err != nil {
			t.Error("should be nil but got error")
		}

		expectedUp := `rename_column("templates", "landing", "html")`
		expectedDown := `rename_column("templates", "html", "landing")`

		if up != expectedUp {
			t.Errorf("expected %v but got %v", expectedUp, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Rename_Index(t *testing.T) {
	r := rename{}

	t.Run("invalid index renamer", func(t *testing.T) {
		_, _, err := r.GenerateFizz("rename_index_idx_from_templates", []string{})
		if err == nil {
			t.Error("should error but got nil")
		}
	})

	t.Run("normal index renamer", func(t *testing.T) {
		up, down, err := r.GenerateFizz("rename_index_idx_events_name_to_idx_events_from_events", []string{})
		if err != nil {
			t.Error("should be nil but got error")
		}

		expectedUp := `rename_index("events", "idx_events_name", "idx_events")`
		expectedDown := `rename_index("events", "idx_events", "idx_events_name")`

		if up != expectedUp {
			t.Errorf("expected %v but got %v", expectedUp, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}

func Test_Rename_Matches(t *testing.T) {
	r := rename{}

	cases := []struct {
		name     string
		expected bool
	}{
		{name: "rena", expected: false},
		{name: "rename", expected: true},
		{name: "rename_table", expected: true},
		{name: "rename_col", expected: true},
		{name: "rename_ind", expected: true},
	}

	for _, c := range cases {
		matchs := r.match(c.name)
		if matchs != c.expected {
			t.Errorf("expected %v but got %v", c.expected, matchs)
		}
	}
}

func Test_Rename_Table_Matches(t *testing.T) {
	r := rename{}

	cases := []struct {
		name     string
		expected error
	}{
		{name: "rename_table", expected: ErrInvalidRenamer},
		{name: "rename_table_users", expected: ErrInvalidRenamer},
		{name: "rename_table_users_to_employees", expected: nil},
	}

	for _, c := range cases {
		_, _, err := r.GenerateFizz(c.name, []string{})
		if !errors.Is(err, c.expected) {
			t.Errorf("expected %v but got %v", c.expected, err)
		}
	}
}

func Test_Rename_Column_Matches(t *testing.T) {
	r := rename{}

	cases := []struct {
		name     string
		expected error
	}{
		{name: "rename_column", expected: ErrInvalidRenamer},
		{name: "rename_column_description", expected: ErrInvalidRenamer},
		{name: "rename_column_description_to_summary", expected: ErrInvalidRenamer},
		{name: "rename_column_description_to_summary_from", expected: ErrInvalidRenamer},
		{name: "rename_column_description_to_summary_from_templates", expected: nil},
	}

	for _, c := range cases {
		_, _, err := r.GenerateFizz(c.name, []string{})
		if !errors.Is(err, c.expected) {
			t.Errorf("expected %v but got %v", c.expected, err)
		}
	}
}
