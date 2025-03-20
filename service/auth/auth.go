package auth

import (
	"fmt"
	"time"

	"github.com/bishalkl/learnBackendApi/config"
	"github.com/golang-jwt/jwt/v4"
)

// Secreat key for signing the JWT
var mySigningKey = []byte(config.Envs.JWT_SECRET)

// Clamis define the structure of the JWT clamis
type Clamis struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenereateJWT generates a new JWT token for the give email
func GenereateJWT(email string) (string, error) {
	// Token expires in 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	Clamis := &Clamis{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "myapp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, Clamis)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}
	return tokenString, nil
}

// Parse JWT parses and validate the JWT toekn from the request
