package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ApplicationDatabase struct {
	db *sql.DB
}

func NewApplicationDatabase() (*ApplicationDatabase, error) {
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

	return &ApplicationDatabase{db: db}, nil
}

func (a *ApplicationDatabase) Close() {
	a.db.Close()
}

func (a *ApplicationDatabase) GetDatabaseVersion() (uint16, error) {
	rows, err := a.db.Query(`SELECT MAX(version) FROM version_history`)
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
	UserID    uint64    `json:"userID"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	BirthDate time.Time `json:"birthDate"`
}

func (a *ApplicationDatabase) ListBabies(userID uint64) ([]*Baby, error) {
	rows, err := a.db.Query(`SELECT baby_id, name, sex, birth_date FROM baby WHERE user_id = ?`, userID)
	if err != nil {
		return nil, fmt.Errorf("error listing babies: %w", err)
	}
	defer rows.Close()

	babies := make([]*Baby, 0)
	for rows.Next() {
		var babyID uint64
		var name string
		var sex string
		var birthDate time.Time

		err = rows.Scan(&babyID, &name, &sex, &birthDate)
		if err != nil {
			return nil, fmt.Errorf("error retrieving baby data: %w", err)
		}

		baby := Baby{Id: babyID, UserID: userID, Name: name, Sex: sex, BirthDate: birthDate}
		babies = append(babies, &baby)
	}

	return babies, nil
}

func (a *ApplicationDatabase) GetBaby(userID uint64, babyID uint64) (*Baby, error) {
	rows, err := a.db.Query(`SELECT name, sex, birth_date FROM baby WHERE user_id = ? AND baby_id = ?`, userID, babyID)
	if err != nil {
		return nil, fmt.Errorf("error getting baby: %w", err)
	}
	defer rows.Close()

	rows.Next()
	var name string
	var sex string
	var birthDate time.Time

	err = rows.Scan(&name, &sex, &birthDate)
	if err != nil {
		return nil, fmt.Errorf("error retrieving baby data: %w", err)
	}

	baby := Baby{Id: babyID, UserID: userID, Name: name, Sex: sex, BirthDate: birthDate}

	return &baby, nil
}

type Event struct {
	Id        uint64     `json:"id"`
	BabyId    uint64     `json:"babyID"`
	EventType string     `json:"eventType"`
	StartTime time.Time  `json:"startTime"`
	EndTime   *time.Time `json:"endTime,omitempty"`
	Notes     *string    `json:"notes,omitempty"`
}

func (a *ApplicationDatabase) ListEvents(userID uint64, babyID uint64) ([]*Event, error) {
	rows, err := a.db.Query(`SELECT e.event_id, e.type, e.start_time, e.end_time, e.notes FROM event e INNER JOIN baby b ON e.baby_id = b.baby_id WHERE b.user_id = ? AND e.baby_id = ?`, userID, babyID)
	if err != nil {
		return nil, fmt.Errorf("error listing events: %w", err)
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		var eventID uint64
		var eventType string
		var startTime time.Time
		var nullableEndTime sql.NullTime
		var nullableNotes sql.NullString

		err = rows.Scan(&eventID, &eventType, &startTime, &nullableEndTime, &nullableNotes)
		if err != nil {
			return nil, fmt.Errorf("error retrieving event data: %w", err)
		}

		var endTime *time.Time
		if nullableEndTime.Valid {
			endTime = &nullableEndTime.Time
		}

		var notes *string
		if nullableNotes.Valid {
			notes = &nullableNotes.String
		}

		event := Event{Id: eventID, BabyId: babyID, EventType: eventType, StartTime: startTime, EndTime: endTime, Notes: notes}
		events = append(events, &event)
	}

	return events, nil
}

func (a *ApplicationDatabase) CreateEvent(babyID uint64, event *Event) (uint64, error) {
	res, err := a.db.Exec(
		`INSERT INTO event (baby_id, type, start_time, end_time, notes) VALUES (?, ?, ?, ?, ?)`,
		babyID, event.EventType, event.StartTime, event.EndTime, event.Notes,
	)
	if err != nil {
		return 0, fmt.Errorf("error creating event: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting id of new event: %w", err)
	}

	return uint64(id), nil
}
