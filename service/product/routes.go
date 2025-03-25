package product

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/bishalkl/learnBackendApi/types"
	"github.com/bishalkl/learnBackendApi/utils"
	"github.com/gorilla/mux"
)

// type
type Handler struct {
	store types.ProductStore
}

// new construtor instance
func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

// for register router
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/product", h.createProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/products", h.GetProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", h.GetProductHandler).Methods(http.MethodGet)
}

// handler for createProduct
func (h *Handler) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product types.Product
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Set the createdAt time
	product.CreatedAt = time.Now()

	// call the store method to create the product
	err := h.store.CreateProduct(&product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}

// hanlder for getProducts
func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with products
	utils.WriteJSON(w, http.StatusOK, products)
}

// handler for getProduct
func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Conver string to int
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.store.GetProductById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, errors.New("Product not found"))
	}

	utils.WriteJSON(w, http.StatusOK, product)

}
