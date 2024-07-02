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

	seed, err := os.ReadFile("database/seed/seed.sql")
	if err != nil {
		return nil, fmt.Errorf("error retrieving database seed file: %w", err)
	}

	_, err = db.Exec(string(seed))
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
