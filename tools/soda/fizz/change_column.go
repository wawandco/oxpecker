package fizz

import (
	"fmt"
	"regexp"
	"strings"
)

var ccReg = regexp.MustCompile(`change_(\w+)_(\w+)`)

type changeColumn struct {
	column     string
	newColType string
	table      string
	isNull     bool
}

func (cc changeColumn) match(name string) bool {
	return ccReg.MatchString(name)
}

func (cc *changeColumn) GenerateFizz(name string, args []string) (string, string, error) {
	if len(args) == 0 {
		return "", "", ErrNoColumnFound
	}

	matches := ccReg.FindAllStringSubmatch(name, -1)[0][1:]
	cc.table = matches[0]
	cc.column = matches[1]
	cc.newColType = columnType(args[0])
	cc.isNull = strings.Contains(args[0], "nulls")

	return cc.fizz(), cc.unFizz(), nil
}

func (cc changeColumn) fizz() string {
	return fmt.Sprintf(`change_column("%s", "%s", "%s", {%s})`, cc.table, cc.column, cc.newColType, cc.parseNull())
}

func (cc changeColumn) unFizz() string {
	return fmt.Sprintf(`change_column("%s", "%s", "string", {})`, cc.table, cc.column)
}

func (cc changeColumn) parseNull() string {
	if !cc.isNull {
		return ""
	}

	return "null: true"
}
