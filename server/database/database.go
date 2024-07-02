package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetApplicationDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:dev.db")
	if err != nil {
		return nil, fmt.Errorf("error connecting to dev database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging dev database: %w", err)
	}

	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("error retrieving database schema: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, fmt.Errorf("error setting up dev database: %w", err)
	}

	_, err = db.Exec(`INSERT OR IGNORE INTO welcome (message) VALUES ('Getting data from the database')`)
	if err != nil {
		return nil, fmt.Errorf("error seeding dev database: %w", err)
	}

	return db, nil
}

func GetWelcomeMessage(db *sql.DB) (string, error) {
	rows, err := db.Query(`SELECT message FROM welcome WHERE id = 1`)
	if err != nil {
		return "", fmt.Errorf("error getting welcome message: %w", err)
	}
	defer rows.Close()

	var name string
	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		return "", fmt.Errorf("error getting welcome message from result: %w", err)
	}
	return name, nil

}
