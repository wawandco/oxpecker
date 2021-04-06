package smartfizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
	"github.com/pkg/errors"
)

type dropTable struct {
	name  string
	table fizz.Table
}

func (dt *dropTable) Generate(args []string) error {
	name := strings.TrimPrefix(dt.name, "drop_table")
	if name == "" {
		return errors.Errorf("no table name")
	}

	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	dt.table = table

	return nil
}

func (dt dropTable) Fizz() string {
	return dt.table.UnFizz()
}

func (dt dropTable) UnFizz() string {
	return dt.table.Fizz()
}
