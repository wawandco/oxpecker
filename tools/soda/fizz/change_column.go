package fizz

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type changeColumn struct {
	column     string
	newColType string
	table      string
}

func (cc changeColumn) match(name string) bool {
	return strings.HasPrefix(name, "change")
}

func (cc *changeColumn) GenerateFizz(name string, args []string) (string, string, error) {
	if len(args) == 0 {
		return "", "", errors.Errorf("no arguments was received, at least 1 column is required")
	}

	reg := regexp.MustCompile(`change_(\w+)_(\w+)`)

	matches := reg.FindAllStringSubmatch(name, -1)[0][1:]
	cc.table = matches[0]
	cc.column = matches[1]
	cc.newColType = cc.parsecolumns(args)

	return cc.fizz(), cc.unFizz(), nil
}

func (cc changeColumn) fizz() string {
	return fmt.Sprintf(`change_column("%s", "%s", "%s", {})`, cc.table, cc.column, cc.newColType)
}

func (cc changeColumn) unFizz() string {
	return fmt.Sprintf(`change_column("%s", "%s", "string", {})`, cc.table, cc.column)
}

func (cc changeColumn) parsecolumns(args []string) string {
	arg := args[0]

	slice := strings.Split(arg, ":")
	if len(slice) == 1 {
		slice = append(slice, "string")
	}

	colType := columnType(slice[1])

	return colType
}
