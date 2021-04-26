package fizz

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gobuffalo/flect"
)

var acReg = regexp.MustCompile(`add(?:ing)?_\w+_to_(\w+)`)

type addColumn struct {
	columns map[string]string
	table   string
}

func (ac addColumn) match(name string) bool {
	return acReg.MatchString(name)
}

func (ac *addColumn) GenerateFizz(name string, args []string) (string, string, error) {
	if len(args) == 0 {
		return "", "", ErrNoColumnFound
	}

	matches := acReg.FindAllStringSubmatch(name, -1)[0][1:]
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
	cols := make(map[string]string)

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
