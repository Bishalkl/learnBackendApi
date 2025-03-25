package types

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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

// CreateProductPayload struct for the product creation request
type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"required,min=10,max=255"`
	Image       string  `json:"image" validate:"required,url"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gt=0"`
}

// Product struct
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=255"`
	Image       string    `json:"image" validate:"required,url"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" validate:"required,gt=0"`
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

// RegisterUserPayload struct for the user registration request
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required,min=3,max=100"`
	LastName  string `json:"lastName" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Token     string `json:"_token,omitempty"`
}

// LoginUserPayload struct for the login request
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Validate validates the RegisterUserPayload struct
func (p *RegisterUserPayload) Validate() error {
	return validate.Struct(p)
}

// Validate validates the LoginUserPayload struct
func (p *LoginUserPayload) Validate() error {
	return validate.Struct(p)
}

// Validate validates the User struct
func (u *User) Validate() error {
	return validate.Struct(u)
}

// Validate validates the product struct
func (p *CreateProductPayload) Validate() error {
	return validate.Struct(p)
}

// Validate validates the Product struct
func (p *Product) Validate() error {
	return validate.Struct(p)
}
