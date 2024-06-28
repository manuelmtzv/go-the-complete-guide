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
		fmt.Println(err)
		panic("could not connect DB.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		panic("could not ping DB.")
	}

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("could not create DB tables.")
	}
}
