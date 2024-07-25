package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create the events table
	err = createTables()

	if err != nil {
		panic(fmt.Sprintf("Could not create tables: %v", err))
	}

}


func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TEXT NOT NULL,
		userID INTEGER
	)
	`

	_, err := DB.Exec(query)

	// Error handling
	if err != nil {
        return fmt.Errorf("error while initializing the database: %v", err)
	}

	return nil
}