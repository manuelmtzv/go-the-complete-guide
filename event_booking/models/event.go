package models

import (
	"event-booking/database"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (event *Event) Save() error {
	query := `
		INSERT INTO 
		events(name, description, location, datetime, user_id) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	var id int64

	err = statement.QueryRow(event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&id)

	if err != nil {
		return err
	}

	event.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events;
	`
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := make([]Event, 0)

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}
