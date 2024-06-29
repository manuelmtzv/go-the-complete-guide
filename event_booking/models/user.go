package models

import (
	"database/sql"
	"event-booking/database"
	"event-booking/utility"
	"fmt"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func GetUser[T int64 | string](identifier T) (*User, error) {
	var query string
	var args []interface{}

	switch v := any(identifier).(type) {
	case int64:
		query = "SELECT id, email, password FROM users WHERE id = $1"
		args = append(args, v)
	case string:
		query = "SELECT id, email, password FROM users WHERE email = $1"
		args = append(args, v)
	default:
		return nil, fmt.Errorf("unsupported identifier type: %T", identifier)
	}

	row := database.DB.QueryRow(query, args...)

	u := User{}
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (u *User) Save() error {
	query := `
		INSERT INTO 
		users(email, password) 
		VALUES ($1, $2) 
		RETURNING id;
	`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	var id int64

	hashedPassword, err := utility.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = statement.QueryRow(u.Email, hashedPassword).Scan(&id)

	u.Id = id

	return err
}
