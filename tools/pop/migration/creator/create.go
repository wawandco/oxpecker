package creator

import "github.com/pkg/errors"

var creations = []Creator{
	&FizzCreator{},
	&SQLCreator{},
}

// CreateMigrationFor selects the correct creation implemention
// Based on the input
func CreateMigrationFor(name string) (Creator, error) {
	for _, creation := range creations {
		if creation.Name() == name {
			return creation, nil
		}
	}

	return nil, errors.Errorf("invalid migration type")
}
