package fizz

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var ErrInvalidRenamer error = errors.New("invalid renamer, please write a valid argument")

type rename struct {
	newName string
	oldName string
	table   string
	renType string
}

func (rt rename) match(name string) bool {
	return strings.HasPrefix(name, "rename")
}

func (rt *rename) GenerateFizz(name string, args []string) (string, string, error) {
	var up, down string
	var reg *regexp.Regexp

	isTable := strings.HasPrefix(name, "rename_table")
	isColumn := strings.HasPrefix(name, "rename_column")
	isIndex := strings.HasPrefix(name, "rename_index")

	if isTable {
		reg = regexp.MustCompile(`rename_table_(\w+)_to_(\w+)`)
		rt.renType = "table"
	}

	if isColumn {
		reg = regexp.MustCompile(`rename_column_(\w+)_to_(\w+)_from_(\w+)`)
		rt.renType = "column"
	}

	if isIndex {
		reg = regexp.MustCompile(`rename_index_(\w+)_to_(\w+)_from_(\w+)`)
		rt.renType = "index"
	}

	if reg == nil {
		return up, down, ErrInvalidRenamer
	}

	if !reg.MatchString(name) {
		return up, down, ErrInvalidRenamer
	}

	matches := reg.FindAllStringSubmatch(name, -1)[0][1:]
	rt.oldName = matches[0]
	rt.newName = matches[1]

	if len(matches) == 3 {
		rt.table = matches[2]
	}

	up = rt.fizz()
	down = rt.unFizz()

	return up, down, nil
}

func (rt rename) fizz() string {
	switch rt.renType {
	case "column":
		return fmt.Sprintf(`rename_column("%s", "%s", "%s")`, rt.table, rt.oldName, rt.newName)
	case "index":
		return fmt.Sprintf(`rename_index("%s", "%s", "%s")`, rt.table, rt.oldName, rt.newName)
	default:
		return fmt.Sprintf(`rename_table("%s", "%s")`, rt.oldName, rt.newName)
	}
}

func (rt rename) unFizz() string {
	switch rt.renType {
	case "column":
		return fmt.Sprintf(`rename_column("%s", "%s", "%s")`, rt.table, rt.newName, rt.oldName)
	case "index":
		return fmt.Sprintf(`rename_index("%s", "%s", "%s")`, rt.table, rt.newName, rt.oldName)
	default:
		return fmt.Sprintf(`rename_table("%s", "%s")`, rt.newName, rt.oldName)
	}
}
