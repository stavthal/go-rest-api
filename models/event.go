package models

import "time"

// Event is a struct that represents an event
type Event struct {
    ID          string    `json:"id"`
    Name      	string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    DateTime    time.Time `json:"dateTime" binding:"required"`
    UserID      string    `json:"userID"`
}

var events = []Event{}



func (e Event) Save() {
	// Save the event to the database
	events = append(events, e)

}

func GetAllEvents() []Event {
	return events
}