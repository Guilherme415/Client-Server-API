package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() error {
	db, err := sql.Open("sqlite3", "clientServerApi.db")
	if err != nil {
		return err
	}

	// err = db.Ping()
	// if err != nil {
	// 	return err
	// }

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}
