package models

import (
	"database/sql"
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

func (e *Event) Save() error {
	query := `
		INSERT INTO 
		events(name, description, location, datetime, user_id) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	var id int64

	err = statement.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&id)

	if err != nil {
		return err
	}

	e.Id = id

	return nil
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

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT id, name, description, location, datetime, user_id
		FROM events
		WHERE id = $1
	`

	row := database.DB.QueryRow(query, id)

	e := Event{}

	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &e, nil
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, datetime = $4
		WHERE id = $5
	`

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	return err
}

func (e *Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.Id)

	return err
}
