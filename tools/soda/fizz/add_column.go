package fizz

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/pkg/errors"
)

type addColumn struct {
	columns map[string]string
	table   string
}

func (ac addColumn) match(name string) bool {
	return strings.HasPrefix(name, "add")
}

func (ac *addColumn) GenerateFizz(name string, args []string) (string, string, error) {
	if len(args) == 0 {
		return "", "", errors.Errorf("no arguments was received, at least 1 column is required")
	}

	reg := regexp.MustCompile(`add_\w+_to_(\w+)`)

	matches := reg.FindAllStringSubmatch(name, -1)[0][1:]
	ac.columns = ac.parsecolumns(args)
	ac.table = matches[0]

	return ac.fizz(), ac.unFizz(), nil
}

func (ac addColumn) fizz() string {
	var cols []string

	for col, val := range ac.columns {
		cols = append(cols, fmt.Sprintf(`add_column("%s", "%s", "%s", {})`, ac.table, col, val))
	}

	return strings.Join(cols, "\n")
}

func (ac addColumn) unFizz() string {
	var cols []string

	for col := range ac.columns {
		cols = append(cols, fmt.Sprintf(`drop_column("%s", "%s")`, ac.table, col))
	}

	return strings.Join(cols, "\n")
}

func (ac *addColumn) parsecolumns(args []string) map[string]string {
	cols := make(map[string]string, 0)
	for _, arg := range args {
		slice := strings.Split(arg, ":")
		if len(slice) == 1 {
			slice = append(slice, "string")
		}

		name := flect.Underscore(slice[0])
		colType := columnType(slice[1])

		cols[name] = colType
	}

	return cols
}
