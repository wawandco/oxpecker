package refresh

import (
	"context"
	"fmt"
	"os"
)

type Initializer struct{}

func (i Initializer) Name() string {
	return "refresh/initializer"
}

func (i *Initializer) Initialize(ctx context.Context, root string, args []string) error {
	// check for database.dev.yml file in root location
	rootYml := root + "/.buffalo.dev.yml"

	content := `app_root: .
	build_target_path : ./cmd/app
	ignored_folders:
	- vendor
	- log
	- logs
	- assets
	- public
	- grifts
	- tmp
	- bin
	- node_modules
	- .sass-cache
	included_extensions:
	- .go
	- .env
	build_path: bin
	build_delay: 200ns
	binary_name: tmp-build
	command_flags: []
	enable_colors: true
	log_name: ox`

	_, err := os.Stat(rootYml)
	if err == nil {

		fmt.Println(".buffalo.dev.yml file already exist")
		return nil

	}
	if os.IsNotExist(err) {

		// create file if it does not exist
		file, err := os.Create(rootYml)

		if err != nil {
			fmt.Println("alo alo")
			return (err)
		}

		_, err = os.OpenFile(rootYml, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("alo alo")
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
