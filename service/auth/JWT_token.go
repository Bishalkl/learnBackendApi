package auth

import (
	"fmt"
	"time"

	"github.com/bishalkl/learnBackendApi/config"
	"github.com/golang-jwt/jwt/v4"
)

// Secret key for signing the JWT
var mySigningKey = []byte(config.Envs.JWT_SECRET)

// Claims defines the structure of the JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for the given email
func GenerateJWT(email string) (string, error) {
	// Token expires in 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    config.Envs.JWTIssure,
		},
	}

	// Use HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}

	return tokenString, nil
}

// ParseJWT parses and validates the JWT token
func ParseJWT(tokenString string) (*Claims, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with the correct method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
