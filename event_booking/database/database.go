package database

import (
	"database/sql"

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
	enableUUIDExtension()

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}

func enableUUIDExtension() {
	enableUuid := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`

	_, err := DB.Exec(enableUuid)

	if err != nil {
		panic(err)
	}
}
