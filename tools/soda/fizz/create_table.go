package fizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/flect"
)

type createTable struct{}

func (ct *createTable) match(name string) bool {
	return strings.HasPrefix(name, "create_table_")
}

func (ct *createTable) GenerateFizz(name string, args []string) (string, string, error) {
	var up, down string
	name = strings.TrimPrefix(name, "create_table_")
	if name == "" {
		return up, down, ErrNoTableName
	}

	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	for _, arg := range args[0:] {
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
