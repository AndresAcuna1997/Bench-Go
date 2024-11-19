package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) OutPutUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "GOD",
			Birthdate: "00/00/0000",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("empty arguments are invalid")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
