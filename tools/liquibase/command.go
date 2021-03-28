package liquibase

import (
	"context"
	"encoding/xml"
	"errors"
	"io/ioutil"

	"github.com/gobuffalo/pop/v5"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/pflag"
	"github.com/wawandco/oxpecker/internal/log"
	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)

var ErrInvalidInstruction = errors.New("Invalid instruction please specify up or down")

type Command struct {
	connectionName string
	steps          int
	connections    map[string]*pop.Connection
	flags          *pflag.FlagSet
}

func (lb Command) Name() string {
	return "migrate"
}

func (lb Command) ParentName() string {
	return "db"
}

func (lb Command) HelpText() string {
	return "runs Liquibase command to update database specified with --conn flag"
}

func (lb *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return lb.Up()
	}

	direction := args[2]
	if direction == "up" {
		return lb.Up()
	}

	if direction == "down" {
		return lb.Rollback()
	}

	return ErrInvalidInstruction
}

func (lb *Command) RunBeforeTest(ctx context.Context, root string, args []string) error {
	lb.connectionName = "test"

	return lb.Up()
}

func (lb Command) Up() error {
	cx := lb.connections[lb.connectionName]
	if cx == nil {
		return errors.New("connection not found")
	}

	conn, err := pgx.Connect(context.Background(), cx.URL())
	if err != nil {
		return err
	}

	err = lb.EnsureTables(conn)
	if err != nil {
		return err
	}

	cl, err := lb.ReadChangelog()
	if err != nil {
		return err
	}

	for _, v := range cl.Migrations {
		// Read the file
		m, err := lb.ReadMigration(v.File)
		if err != nil {
			return err
		}

		for _, mc := range m.ChangeSets {
			err = mc.Execute(conn, v.File)
			if err != nil {
				log.Errorf("error executing `%v`.", mc.ID)
				return err
			}
		}
	}

	log.Info("Database up to date.")

	return nil
}

func (lb *Command) Rollback() error {
	cx := lb.connections[lb.connectionName]
	if cx == nil {
		return errors.New("connection not found")
	}

	conn, err := pgx.Connect(context.Background(), cx.URL())
	if err != nil {
		return err
	}

	err = lb.EnsureTables(conn)
	if err != nil {
		return err
	}

	// Default to 1 on down.
	if lb.steps == 0 {
		lb.steps = 1
	}

	for i := 0; i < lb.steps; i++ {
		var id, file string
		row := conn.QueryRow(context.Background(), `SELECT filename, id FROM databasechangelog ORDER BY orderexecuted desc`)
		err = row.Scan(&file, &id)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		if errors.Is(err, pgx.ErrNoRows) {
			log.Info("no migrations to run down.")

			return nil
		}

		m, err := lb.ReadMigration(file)
		if err != nil {
			return err
		}

		for _, v := range m.ChangeSets {
			if v.ID != id {
				continue
			}

			err := v.Rollback(conn)
			if err != nil {
				log.Errorf("error rolling back `%v`.\n", v.ID)

				return err
			}
		}
	}

	return nil
}

func (lb *Command) ParseFlags(args []string) {
	lb.flags = pflag.NewFlagSet(lb.Name(), pflag.ContinueOnError)
	lb.flags.StringVarP(&lb.connectionName, "conn", "", "development", "the name of the connection to use")
	lb.flags.IntVarP(&lb.steps, "steps", "s", 0, "number of migrations to run")
	lb.flags.Parse(args) //nolint:errcheck,we don't care hence the flag
}

func (lb *Command) Flags() *pflag.FlagSet {
	return lb.flags
}

func (lb Command) ReadChangelog() (*ChangeLog, error) {
	d, err := ioutil.ReadFile("migrations/changelog.xml")
	if err != nil {
		return nil, err
	}

	cl := &ChangeLog{}
	err = xml.Unmarshal([]byte(d), cl)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (lb Command) ReadMigration(path string) (*Migration, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	m := &Migration{}
	err = xml.Unmarshal([]byte(d), m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
