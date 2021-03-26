package flect

import (
	"context"
	"fmt"
	"os"

	"github.com/wawandco/oxpecker/lifecycle/new"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "flect/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, options new.Options) error {

	yml := options.Folder + "/inflections.yml"
	content := `{ "singular": "plural" }`

	_, err := os.Stat(yml)
	if err == nil {
		fmt.Println("inflections.yml file already exist ")

		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}
	// create file if it does not exist
	file, err := os.Create(yml)
	if err != nil {
		return (err)
	}

	_, err = os.OpenFile(yml, os.O_RDWR, 0644)
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
