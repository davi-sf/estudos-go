package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO t_events(name, description, location, dateTime, user_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM t_events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM t_events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) Update() error {
	query := `
	UPDATE t_events
	SET name = $1, description = $2, location = $3, dateTime = $4, user_id = $5
	WHERE id = $6`

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.Id)
	return err
}

func DeleteEventByID(id int64) error {
	query := `DELETE FROM t_events WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO t_registrations(event_id, user_id) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, e.Id, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM t_registrations WHERE event_id = $1 AND user_id = $2"
	_, err := db.DB.Exec(query, e.Id, userId)
	return err
}
