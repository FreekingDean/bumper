package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDatabasePath = "/var/lib/bumper/db.sqlite"
)

type service struct {
	db *sql.DB

	config configuration
}

type configuration interface {
	AddOption(string, interface{})
	GetString(string) String
}

// Init tells the configuration engine how
// to configure the DB
func Init(config configuration) *service {
	svc := &service{
		config: config,
	}
	config.AddOption("path", defaultDatabasePath)
	return svc
}

func (svc *service) Start() error {
	db, err := sql.Open("sqlite3", svc.config.GetString("path"))
	if err != nil {
		return err
	}
	svc.db = db
	return nil
}

func (d *Database) Stop() {
	d.Close()
}

func (d *Database) Store(statement string, data ...interface{}) error {
	stmt, err := d.db.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data...)
	if err != nil {
		return err
	}
	return nil
}
