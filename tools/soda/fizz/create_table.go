package fizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

type createTable struct{}

func (ct *createTable) match(name string) bool {
	return strings.HasPrefix(name, "create_table")
}

func (ct *createTable) GenerateFizz(name string, args []string) (string, string, error) {
	var up, down string
	name = strings.TrimPrefix(name, "create_table")
	if name == "" {
		return up, down, errors.Errorf("no table name")
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
			return up, down, err
		}
	}

	up = table.Fizz()
	down = table.UnFizz()

	return up, down, nil
}
