package models

import (
	"database/sql"
	"example/rest_api/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"dateTime"`
	UserID      int       `json:"user_id"`
}

var events []Event

func (e Event) Save() error {
	// save the event to the database

	query := ` INSERT INTO events(name, description, location, dateTime, user_id)
	 VALUES(?, ?, ?, ?, ?)`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(prepare)
	result, err := prepare.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(rows)

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
