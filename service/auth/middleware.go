package auth

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

// Logging Middleware logs the request method, URL path, and the time it took to the process the request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request details
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, time.Since(start))

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// JWTMiddlware checks if the request contains a vaild JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bear <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		claims, err := ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user email to request context for use in other handlers
		ctx := r.Context() // Create a new context to store the user information
		ctx = context.WithValue(ctx, "user-email", claims.Email)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
