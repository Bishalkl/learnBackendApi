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

}

// handler for register
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// getJSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if the user exists
	existingUser, err := h.store.GetUserByEmail(payload.Email)
	if err == nil && existingUser != nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("User already exists"))
		return
	}

	// Hash password
	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
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
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Return success response
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User register successfully"})

}
