package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bishalkl/learnBackendApi/service/auth"
	"github.com/bishalkl/learnBackendApi/types"
	"github.com/bishalkl/learnBackendApi/utils"
	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	store types.UserStore
}

// creating instance of Handler
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRouter sets up routes
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handler for login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Parse the login payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request payload"))
		return
	}

	// Check if the user exists
	existingUser, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid email or password"))
		return
	}

	// Check if the provided password is correct
	if !auth.ComparePassword(existingUser.Password, payload.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid credentials"))
		return
	}

	// Generate JWT token (assuming you have functin for this)
	token, err := auth.GenerateJWT(existingUser.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Faild to generate token"))
		return
	}

	// Send success response with JWT token
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   token,
	})

}

// handler for register
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// GetJSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request payload"))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("User already exists"))
		return
	}

	// Hash password
	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to hash password"))
		return
	}

	// Create a new user
	newUser := &types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	// Save user in DB
	err = h.store.CreateUser(newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to create user"))
		return
	}

	// Return success response
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User register successfully"})

}
