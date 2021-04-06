package smartfizz

import (
	"strings"
)

// SmartFizzer follows a set of steps to make possible the creation
type SmartFizzer interface {
	Generate([]string) error
	Fizz() string
	UnFizz() string
}

func New(name string) SmartFizzer {
	sf := &createTable{name: name}

	if strings.HasPrefix(name, "create_table") {
		return &createTable{name: name}
	}

	if strings.HasPrefix(name, "drop_table") {
		return &dropTable{name: name}
	}

	return sf
}
