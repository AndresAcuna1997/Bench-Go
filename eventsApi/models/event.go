package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
    VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		fmt.Println("Error executing statement:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return err
	}
	e.ID = id
	return nil
}

func (e *Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, dateTime =?
		WHERE id = ?
		`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error updating event:", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM events WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func GetAllEvent() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventFromId(eventID int64) (Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, eventID)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		fmt.Println("Error scanning row:", err)
		return Event{}, err
	}

	return event, nil
}
