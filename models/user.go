package models

import (
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

// User struct
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	// Insert the user into the database
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	// Prepare the query
	stmt, err := db.DB.Prepare(query)

	// Error handling
	if err != nil {
		return fmt.Errorf("error preparing the query: %v", err)
	}


	defer stmt.Close()

	// Hash the password
	hashedPassword, err := utils.HashPassword(u.Password)

	// Error handling
	if err != nil {
		return fmt.Errorf("error hashing the password: %v", err)
	}


	// Execute the query
	_, err = stmt.Exec(u.Email, hashedPassword)


	// Error handling
	if err != nil {
		return fmt.Errorf("error saving the user: %v", err)
	}	

	return nil
}

func (u User) Authenticate() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	// Get the user ID and password from the row that was found from the query
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.ComparePasswords(u.Password, retrievedPassword)

	if !passwordIsValid {
		return fmt.Errorf("invalid password")
	}

	return nil
}

