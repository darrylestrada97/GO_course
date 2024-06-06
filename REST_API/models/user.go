package models

import (
	"database/sql"
	"example/rest_api/db"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	// save the user to the database

	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		prepare.Close()
	}(prepare)

	result, err := prepare.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = int(userId)
	return err
}
