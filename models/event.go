package models

import (
	"example.com/rest-api/db"
)

// Event is a struct that represents an event
type Event struct {
    ID          string    `json:"id"`
    Name      	string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    DateTime    string    `json:"dateTime" binding:"required"`
    UserID      string    `json:"userID"`
}




func (e Event) Save() error {
	// Save the event to the database
    query := `
    INSERT INTO events (name, description, location, dateTime, userID)
    VALUES (?, ?, ?, ?, ?)
    `

    // Prepare the query
    stmt, prepareError := db.DB.Prepare(query)

    if prepareError != nil {
        panic(prepareError)
    }

    // Execute the query
    _, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)


    defer stmt.Close() // Make sure to close the statement when the function returns

    if err != nil {
        return err
    }

    return nil

}

func GetAllEvents() ([]Event, error) {
    // Query the database for all events
    query := `
    SELECT * FROM events
    `

    rows, err := db.DB.Query(query)

    if err != nil {
        panic(err)
    }

    defer rows.Close() // Make sure to close the rows when the function returns

    var events []Event

    for rows.Next() {
        var event Event

        err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

        if err != nil {
            return nil, err
        }

        events = append(events, event)
    }

    return events, nil
}