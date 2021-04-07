package fizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

type createTable struct {
	name  string
	table fizz.Table
}

func (ct *createTable) Generate(args []string) error {
	name := strings.TrimPrefix(ct.name, "create_table")
	if name == "" {
		return errors.Errorf("no table name")
	}

	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	for _, arg := range args[1:] {
		slice := strings.Split(arg, ":")
		if len(slice) == 1 {
			slice = append(slice, "string")
		}

		o := fizz.Options{}
		name := flect.Underscore(slice[0])
		colType := columnType(slice[1])

		if name == "id" {
			o["primary"] = true
		}

		if strings.HasPrefix(strings.ToLower(slice[1]), "nulls.") {
			o["null"] = true
		}

		if err := table.Column(name, colType, o); err != nil {
			return err
		}
	}

	ct.table = table

	return nil
}

func (ct createTable) Fizz() string {
	return ct.table.Fizz()
}

func (ct createTable) UnFizz() string {
	return ct.table.UnFizz()
}
