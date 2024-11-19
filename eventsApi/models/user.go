package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.Id = userID
	return err
}

func (u *User) ValidateCredential() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string

	err := row.Scan(&u.Id, &retrievedPass)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPass)

	if !passwordIsValid {
		return errors.New("credetials invalid")
	}

	return nil
}

func ListUsers() ([]User, error) {
	query := "SELECT * FROM users"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
