package flect

import (
	"context"
	"fmt"
	"os"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "flect/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {

	rootYml := root + "/inflections.yml"

	content := `
	{
	  "singular": "plural"
	}
	`

	_, err := os.Stat(rootYml)
	if err == nil {

		fmt.Println("inflections.yml file already exist ")
		return nil

	}
	if os.IsNotExist(err) {

		// create file if it does not exist
		file, err := os.Create(rootYml)

		if err != nil {
			return (err)
		}

		_, err = os.OpenFile(rootYml, os.O_RDWR, 0644)
		if err != nil {
			return (err)
		}

		_, err = file.WriteString(content)
		if err != nil {
			return (err)
		}

		file.Close()

		return nil

	}

	return err

}
