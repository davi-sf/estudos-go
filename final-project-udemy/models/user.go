package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO t_users(email, password)
	VALUES ($1, $2) RETURNING id`

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, user.Email, hashedPassword).Scan(&user.Id)
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM t_users WHERE email = $1"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
