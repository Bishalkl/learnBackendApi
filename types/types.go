package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductById(UserID int) (Product, error)
}

// types for Product
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"_token,omitempty"`
}

// types LoginUserPayload struct
type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
