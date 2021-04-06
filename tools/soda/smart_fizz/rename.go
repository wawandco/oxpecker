package smartfizz

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type rename struct {
	name    string
	newName string
	oldName string
	table   string
	renType string
}

func (rt *rename) Generate(args []string) error {
	var reg *regexp.Regexp

	isTable := strings.HasPrefix(rt.name, "rename_table")
	isColumn := strings.HasPrefix(rt.name, "rename_column")
	isIndex := strings.HasPrefix(rt.name, "rename_index")

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

	if !reg.MatchString(rt.name) {
		return errors.Errorf("invalid renamer, please write a valid argument")
	}

	matches := reg.FindAllStringSubmatch(rt.name, -1)[0][1:]
	rt.oldName = matches[0]
	rt.newName = matches[1]

	if len(matches) == 3 {
		rt.table = matches[2]
	}

	return nil
}

func (rt rename) Fizz() string {
	switch rt.renType {
	case "column":
		return fmt.Sprintf(`rename_column("%s", "%s", "%s")`, rt.table, rt.oldName, rt.newName)
	case "index":
		return fmt.Sprintf(`rename_index("%s", "%s", "%s")`, rt.table, rt.oldName, rt.newName)
	default:
		return fmt.Sprintf(`rename_table("%s", "%s")`, rt.oldName, rt.newName)
	}
}

func (rt rename) UnFizz() string {
	switch rt.renType {
	case "column":
		return fmt.Sprintf(`rename_column("%s", "%s", "%s")`, rt.table, rt.newName, rt.oldName)
	case "index":
		return fmt.Sprintf(`rename_index("%s", "%s", "%s")`, rt.table, rt.newName, rt.oldName)
	default:
		return fmt.Sprintf(`rename_table("%s", "%s")`, rt.newName, rt.oldName)
	}
}
