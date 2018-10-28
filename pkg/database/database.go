package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DATABASE_PATH = "./bumper.sqlite"
)

type Database struct {
	db *sql.DB
}

func Open() (*Database, error) {
	db, err := sql.Open("sqlite3", DATABASE_PATH)
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
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
