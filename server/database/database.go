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

func GetDatabaseVersion(db *sql.DB) (uint16, error) {
	rows, err := db.Query(`SELECT MAX(version) FROM version_history`)
	if err != nil {
		return 0, fmt.Errorf("error getting db version: %w", err)
	}
	defer rows.Close()

	var version uint16
	rows.Next()
	err = rows.Scan(&version)
	if err != nil {
		return 0, fmt.Errorf("error getting db version from result: %w", err)
	}
	return version, nil
}
