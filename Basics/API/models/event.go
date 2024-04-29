package models

import (
	"fmt"
	"time"

	"example.com/api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// INSERT
func (e Event) Save() error {
	//later: add it to a database
	query := "INSERT INTO events (name, description, location, dateTime, user_id) values (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return (err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return (err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return (err)
	}

	e.ID = id
	return err
}

// SELECT
func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		var dateTimeStr string // Use uma string temporária para verificar o dado bruto.
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)

		if err != nil {
			return nil, err
		}

		event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
		if err != nil {
			fmt.Println("Error parsing dateTime:", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
