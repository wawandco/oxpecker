package migrate

import "github.com/spf13/pflag"

func (m *Plugin) ParseFlags(args []string) {
	m.flags = pflag.NewFlagSet(m.Name(), pflag.ContinueOnError)

	m.flags.StringVarP(&m.connectionName, "conn", "", "development", "the name of the connection to use")
	m.flags.StringVarP(&m.migrationPath, "folder", "", "./migrations", "the path to the migrations")
	m.flags.StringVarP(&m.direction, "direction", "", "", "direction to run the migrations to")
	m.flags.StringVarP(&m.configFile, "config", "", "config/database.yml", "direction to run the migrations to")
	m.flags.IntVarP(&m.steps, "steps", "s", 0, "how many migrations to run")
	m.flags.Parse(args)
}

func (m *Plugin) Flags() *pflag.FlagSet {
	return m.flags
}
