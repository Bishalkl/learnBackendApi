package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Email     string `json: "email"`
	Password  string `json: "password"`
	Token     string `json: "_token,omitempty"`
}

// types LoginUserPayload struct
type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
