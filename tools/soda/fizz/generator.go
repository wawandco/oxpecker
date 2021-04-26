package fizz

import "errors"

var generators = MigrationGenerators{
	&addColumn{},
	&changeColumn{},
	&createTable{},
	&dropTable{},
	&rename{},
	&dropIndex{},
}

// Errors
var (
	ErrExpressionNotMatch error = errors.New("generator do not match a valid expression")
	ErrNoColumnFound      error = errors.New("no arguments was received, at least 1 column is required")
	ErrNoTableName        error = errors.New("no table name")
)

type MigrationGenerator interface {
	match(string) bool
	GenerateFizz(string, []string) (string, string, error)
}

type MigrationGenerators []MigrationGenerator

func (a MigrationGenerators) GeneratorFor(name string) MigrationGenerator {
	// Setting create table migration by default
	var mg MigrationGenerator = &createTable{}

	for _, x := range a {
		if x.match(name) {
			mg = x
			break
		}
	}

	return mg
}
