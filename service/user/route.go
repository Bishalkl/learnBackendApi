package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
}

// creating instance of Handler
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRouter sets up routes
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handler for login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handler for register
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Registration successfull"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
