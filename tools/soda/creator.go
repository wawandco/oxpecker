package soda

type Creator interface {
	// Receives a migration type string and returns if it
	// applies to it or not.
	Creates(string) bool

	// Creates the migration.
	Create(dir, name string, args []string) error
}

type Creators []Creator

func (a Creators) CreatorFor(name string) Creator {
	for _, x := range a {
		if x.Creates(name) {
			return x
		}
	}

	return nil
}
