package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string       `json:"id"`
	Created   time.Time    `json:"created"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	LastLogin sql.NullTime `json:"lastLogin,omitempty"`
}

func NewUser(id string, created time.Time, username string, email string, password string) User {
	return User{
		Id:       id,
		Created:  created,
		Username: username,
		Email:    email,
		Password: password,
	}
}
