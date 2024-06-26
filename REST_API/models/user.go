package models

import (
	"database/sql"
	"example/rest_api/db"
	"example/rest_api/utils"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	// save the user to the database

	query := " INSERT INTO users(email, password) VALUES(?,?)"
	prepare, err := db.DB.Prepare(query)
	if err != nil {

		return err
	}
	defer func(prepare *sql.Stmt) {
		_ = prepare.Close()
	}(prepare)

	hashedPassword, err := utils.HashPassword(u.Password)
	result, err := prepare.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Print(err)
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = int(userId)
	return err
}

func (u *User) VerifyUser() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, password)

	if !passwordIsValid {
		return fmt.Errorf("invalid password")
	}
	return nil
}
