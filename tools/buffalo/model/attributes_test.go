package model

import (
	"reflect"
	"testing"

	"github.com/gobuffalo/flect/name"
)

func Test_BuildAttrs(t *testing.T) {
	defaults := []attr{{Name: name.New("id"), CommonType: "uuid"}, {Name: name.New("created_at"), CommonType: "timestamp"}, {Name: name.New("updated_at"), CommonType: "timestamp"}}

	cases := []struct {
		args     []string
		expected []attr
		testName string
	}{
		{
			testName: "Empty Args",
			args:     []string{},
			expected: defaults,
		},
		{
			testName: "Some Args Without Type",
			args:     []string{"description:text", "title"},
			expected: []attr{{Name: name.New("id"), CommonType: "uuid"}, {Name: name.New("created_at"), CommonType: "timestamp"}, {Name: name.New("updated_at"), CommonType: "timestamp"}, {Name: name.New("description"), CommonType: "text"}, {Name: name.New("title"), CommonType: "string"}},
		},
		{
			testName: "Replacing Defaults",
			args:     []string{"description:text", "id:int"},
			expected: []attr{{Name: name.New("created_at"), CommonType: "timestamp"}, {Name: name.New("updated_at"), CommonType: "timestamp"}, {Name: name.New("description"), CommonType: "text"}, {Name: name.New("id"), CommonType: "int"}},
		},
		{
			testName: "Replacing Defaults 2",
			args:     []string{"created_at:int", "description:text", "updated_at:int", "id:int"},
			expected: []attr{{Name: name.New("created_at"), CommonType: "int"}, {Name: name.New("description"), CommonType: "text"}, {Name: name.New("updated_at"), CommonType: "int"}, {Name: name.New("id"), CommonType: "int"}},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			attrs := buildAttrs(c.args)
			if !reflect.DeepEqual(c.expected, attrs) {
				t.Errorf("unexpected result, it should be %v but got %v", c.expected, attrs)
			}
		})
	}
}
