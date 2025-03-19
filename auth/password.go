package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// hashpassword
func HashPassword(password string) (string, error) {
	// Generate a hashed password with a cost of 14 (higher cost = most security, but slower)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
