package liquibase

import "encoding/xml"

// DatabaseChangelog file containing all of the migrations.
// This is the root migrations commander, the tool only considers
// migration files in the changelog.
type ChangeLog struct {
	XMLName    xml.Name        `xml:"databaseChangeLog"`
	Migrations []MigrationFile `xml:"include"`
}

// MigrationFiles in the changelog. These are used to
// get the file path to the migration.
type MigrationFile struct {
	File string `xml:"file,attr"`
}
