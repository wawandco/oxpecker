package model

import (
	"reflect"
	"testing"

	"github.com/gobuffalo/flect/name"
)

func Test_BuildImports(t *testing.T) {
	cases := []struct {
		attrs    []attr
		expected []string
		testName string
	}{
		{
			testName: "With Default Attributes",
			attrs:    []attr{{Name: name.New("id"), CommonType: "uuid"}, {Name: name.New("created_at"), CommonType: "timestamp"}, {Name: name.New("updated_at"), CommonType: "timestamp"}},
			expected: []string{"fmt", "github.com/gofrs/uuid", "time"},
		},
		{
			testName: "All Possible Attributes",
			attrs:    []attr{{Name: name.New("id"), CommonType: "uuid"}, {Name: name.New("created_at"), CommonType: "timestamp"}, {Name: name.New("updated_at"), CommonType: "timestamp"}, {Name: name.New("description"), CommonType: "nulls.String"}, {Name: name.New("prices"), CommonType: "slices.Float"}},
			expected: []string{"fmt", "github.com/gobuffalo/nulls", "github.com/gobuffalo/pop/v5/slices", "github.com/gofrs/uuid", "time"},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			imports := buildImports(c.attrs)
			if !reflect.DeepEqual(c.expected, imports) {
				t.Errorf("unexpected result, it should be %v but got %v", c.expected, imports)
			}
		})
	}
}
