package database

import (
	"api/src/config"
	"database/sql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
