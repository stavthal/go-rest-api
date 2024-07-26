package models

import (
	"example.com/rest-api/db"
)

// Event is a struct that represents an event
type Event struct {
    ID          int64     `json:"id"`
    Name      	string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    DateTime    string    `json:"dateTime" binding:"required"`
    UserID      int64     `json:"userID"`
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


    defer stmt.Close() // Closing the stmt here makes the use of stmt pointless. It should not be closed to take advantage of the prepared statement

    if err != nil {
        return err
    }

    return nil

}

func (e Event) Delete() error {
    // Delete the event from the database
    query := `
    DELETE FROM events
    WHERE id = ?
    `

    // Prepare the query
    stmt, prepareError := db.DB.Prepare(query)

    if prepareError != nil {
        panic(prepareError)
    }

    // Execute the query
    _, err := stmt.Exec(e.ID)

    defer stmt.Close()

    if err != nil {
        return err
    }

    return nil
}

func (e Event) Update(id int64) error {
    // Update the event in the database
    query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, dateTime = ?
    WHERE id = ?
    `

    // Prepare the query
    stmt, prepareError := db.DB.Prepare(query)

    if prepareError != nil {
        panic(prepareError)
    }

    // Execute the query
    _, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, id)

    defer stmt.Close()

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

func GetEventById (id int64) (Event, error) {
    // Query the database for an event by id
    query := `
    SELECT * FROM events WHERE id = ?
    `

    var event Event

    err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

    if err != nil {
        return Event{
            ID: -1,
        }, err
    }

    return event, nil
}

