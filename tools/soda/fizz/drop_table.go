package fizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
)

type dropTable struct{}

func (dt dropTable) match(name string) bool {
	return strings.HasPrefix(name, "drop_table_")
}

func (dt *dropTable) GenerateFizz(name string, args []string) (string, string, error) {
	var up, down string
	name = strings.TrimPrefix(name, "drop_table_")
	if name == "" {
		return up, down, ErrNoTableName
	}

	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	return table.UnFizz(), table.Fizz(), nil
}
