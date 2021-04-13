package fizz

var generators = MigrationGenerators{
	&addColumn{},
	&createTable{},
	&dropTable{},
	&rename{},
}

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
