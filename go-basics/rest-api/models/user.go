package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/models/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "insert into users (email, password) values (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, pass)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
