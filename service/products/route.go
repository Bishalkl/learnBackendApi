package products

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// func for registerRouter
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/products", h.getProducts).Methods("GET")
}

// handler for get all products
func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {

}
