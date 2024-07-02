package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const setup = `
-- CREATE DATABASE babystats;

CREATE TABLE IF NOT EXISTS welcome (
  id INT PRIMARY KEY NOT NULL,
  message TEXT NOT NULL
);

INSERT OR IGNORE INTO welcome (id, message) VALUES (1, 'Getting data from the database');
`

func GetApplicationDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:dev.db")
	if err != nil {
		return nil, fmt.Errorf("error connecting to dev database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging dev database: %w", err)
	}

	_, err = db.Exec(setup)
	if err != nil {
		return nil, fmt.Errorf("error setting up dev database: %w", err)
	}

	return db, nil
}

func GetWelcomeMessage(db *sql.DB) (string, error) {
	rows, err := db.Query(`SELECT message FROM welcome WHERE id = 1`)
	if err != nil {
		return "", fmt.Errorf("error getting welcome message: %w", err)
	}

	var name string
	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		return "", fmt.Errorf("error getting welcome message from result: %w", err)
	}
	return name, nil

}
