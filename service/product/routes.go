package product

import (
	"fmt"
	"net/http"

	"github.com/bishalkl/learnBackendApi/types"
	"github.com/gorilla/mux"
)

// type
type Handler struct {
	store types.ProductStore
}

// new construtor instance
func NewHanlder(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

// for register router
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/product", h.getProduct).Methods("GET")
}

// hanlder for getProduct
func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,this is getProduct page")
}
