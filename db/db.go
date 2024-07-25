package db

import (
	"database/sql"

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
	createTables()

}


func createTables() {
	usersQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL unique,
		password TEXT NOT NULL
	)`

	// Execute the query for users
	_, userErr := DB.Exec(usersQuery)

	// Error handling
	if userErr != nil {
    	panic("Count not create the users table")
	}


	eventsQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TEXT NOT NULL,
		userID INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	)`

	// Execute the query for events
	_, err := DB.Exec(eventsQuery)

	// Error handling
	if err != nil {
		panic("Count not create the events table")
	}

}