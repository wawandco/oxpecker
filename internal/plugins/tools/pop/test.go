package pop

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobuffalo/pop/v5"
)

func (p *Plugin) RunBeforeTest(ctx context.Context, root string, args []string) error {
	_, err := os.Stat(filepath.Join(root, "config", "database.yml"))
	if err != nil {
		return err
	}

	db, err := pop.Connect("test")
	if err != nil {
		return err
	}

	fmt.Println(">>> Resetting Database")
	err = db.Dialect.DropDB()
	if err != nil {
		fmt.Printf("could not drop `%v` database, continuing.", db.Dialect.Name())
	}

	err = db.Dialect.CreateDB()
	if err != nil {
		return err
	}

	// Running migrations
	fmt.Println(">>> Running migrations")
	ms := filepath.Join(root, "migrations")
	fm, err := pop.NewFileMigrator(ms, db)
	if err != nil {
		return err
	}

	return fm.Up()
}
