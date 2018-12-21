package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDatabasePath = "/var/lib/bumper/db.sqlite"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	d.db.Close()
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
