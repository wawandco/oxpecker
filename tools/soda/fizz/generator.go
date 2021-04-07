package fizz

type MigrationGenerator interface {
	AppliesTo(string) bool
	GenerateFizz(string, []string) error
}

type MigrationGenerators []MigrationGenerator

func (a MigrationGenerators) GeneratorFor(name string) *MigrationGenerator {
	for _, x := range a {
		if x.AppliesTo(name) {
			return &x
		}
	}

	return nil
}
