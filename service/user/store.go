package user

import (
	"database/sql"
	"fmt"

	"github.com/bishalkl/learnBackendApi/types"
)

// Store interacts with the database
type Store struct {
	db *sql.DB
}

// NewStore creates a new instance of Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserByEmail fetches a user by email
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	row := s.db.QueryRow("SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?", email)

	u := new(types.User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return u, nil
}

// CreateUser for save to database
func (s *Store) CreateUser(user *types.User) error {
	// SQL query to insert user data
	query := `INSERT INTO users (firtName, lastName, email, password, createdAt) VALUES(?, ?, ?, ?, ?)`

	// Execute the query and insert the user
	_, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
