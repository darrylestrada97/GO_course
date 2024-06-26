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

func (e *Event) Save() error {
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
	var events []Event
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

func GetEvent(id string) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {

	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(prepare *sql.Stmt) {
		_ = prepare.Close()
	}(prepare)

	_, err = prepare.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		_ = prepare.Close()
	}(prepare)
	_, err = prepare.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Register(userId int) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?, ?)`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		_ = prepare.Close()
	}(prepare)
	_, err = prepare.Exec(e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(userId int) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		_ = prepare.Close()
	}(prepare)
	_, err = prepare.Exec(e.ID, userId)
	return err
}
