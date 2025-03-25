package types

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductById(id int) (*Product, error)
	CreateProduct(product *Product) error
}

// for validate
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// types for Product
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required, min=3, max=100"`
	Description string    `json:"description" validate:"required, min=10, max=255"`
	Image       string    `json:"image" validate:"required, url"`
	Price       float64   `json:"price" validate:"required, gt=0"`
	Quantity    int       `json:"quantity" validate:"required, gt=0"`
	CreatedAt   time.Time `json:"createdAt"`
}

// User struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName" validate:"required,min=3,max=100"`
	LastName  string    `json:"lastName" validate:"required,min=3,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=6"`
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
