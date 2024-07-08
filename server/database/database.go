package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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

type Baby struct {
	Id        uint64    `json:"id"`
	UserID    uint64    `json:"userId"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
}

func ListBabies(db *sql.DB, userID uint64) ([]Baby, error) {
	rows, err := db.Query(`SELECT baby_id, user_id, name, birth_date FROM baby WHERE user_id = ?`, userID)
	if err != nil {
		return nil, fmt.Errorf("error listing babies: %w", err)
	}
	defer rows.Close()

	babies := make([]Baby, 0)
	for rows.Next() {
		var babyID uint64
		var userID uint64
		var name string
		var birthDate time.Time

		err = rows.Scan(&babyID, &userID, &name, &birthDate)
		if err != nil {
			return nil, fmt.Errorf("error retrieving baby data: %w", err)
		}

		baby := Baby{Id: babyID, UserID: userID, Name: name, BirthDate: birthDate}
		babies = append(babies, baby)
	}

	return babies, nil
}
