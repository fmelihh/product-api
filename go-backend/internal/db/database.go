package db

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDatabase(dbPath string) error {
	var err error
	if db, err = sql.Open("sqlite", dbPath); err != nil {
		return err
	}
	db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS product (
				id TEXT PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				price FLOAT,
				description TEXT,
				currency TEXT,
				category TEXT,
				location TEXT,
				created_at DATETIME,
				updated_at DATETIME
		)`,
	)
	log.Println("PRODUCT TABLE WAS CREATED")
	return nil
}
