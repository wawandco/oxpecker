package db

type Resetter interface {
	DropDB() error
	CreateDB() error
}
