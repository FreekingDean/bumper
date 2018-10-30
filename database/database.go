package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	db *sql.DB
}

type configuration interface {
	BindOption(string, interface{}, interface{})
}

// Init tells the configuration engine how
// to configure the DB
func Init(config configuration) *service {
	svc := &service{}
	config.AddOption("path", DATABASE_PATH)
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
