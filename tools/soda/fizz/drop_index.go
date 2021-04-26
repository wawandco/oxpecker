package fizz

import (
	"fmt"
	"regexp"
)

var diReg = regexp.MustCompile(`drop_index_(\w+)_from_(\w+)`)

type dropIndex struct {
	index string
	table string
}

func (di dropIndex) match(name string) bool {
	return diReg.MatchString(name)
}

func (di *dropIndex) GenerateFizz(name string, args []string) (string, string, error) {
	matches := diReg.FindAllStringSubmatch(name, -1)[0][1:]
	di.index = matches[0]
	di.table = matches[1]

	return di.fizz(), di.unFizz(), nil
}

func (di dropIndex) fizz() string {
	return fmt.Sprintf(`drop_index("%s", "%s")`, di.table, di.index)
}

func (di dropIndex) unFizz() string {
	return fmt.Sprintf(`add_index("%s", "%s", {})`, di.table, di.index)
}
