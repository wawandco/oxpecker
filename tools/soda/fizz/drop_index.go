package fizz

import (
	"fmt"
	"regexp"
	"strings"
)

type dropIndex struct {
	index string
	table string
}

func (di dropIndex) match(name string) bool {
	return strings.HasPrefix(name, "drop_index")
}

func (di *dropIndex) GenerateFizz(name string, args []string) (string, string, error) {
	reg := regexp.MustCompile(`drop_index_(\w+)_from_(\w+)`)

	matches := reg.FindAllStringSubmatch(name, -1)[0][1:]
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
