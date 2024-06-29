package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDatabase() {
	connectionString := "postgres://postgres:password@localhost:5432/event-booking-db"

	var err error
	DB, err = sql.Open("pgx", connectionString)

	if err != nil {
		panic("could not connect DB.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		panic("could not ping DB.")
	}

	createTables()
}

func createTables() {
	createUsersTable()
	createEventsTable()
}

func createUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, 
		email TEXT NOT NULL UNIQUE, 
		password TEXT NOT NULL
	);
	`

	createTable(query, "users")
}

func createEventsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP,
		user_id INTEGER, 
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	createTable(query, "events")
}

func createTable(query string, tableName string) {
	_, err := DB.Exec(query)

	if err != nil {
		fmt.Printf("Could not create %v table.\n", tableName)
		panic(err)
	}
}
