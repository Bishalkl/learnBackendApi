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

// Handler type to handle product-related operations
type Handler struct {
	store types.ProductStore
}

// NewHandler creates a new Handler instance
func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRouter registers the product-related routes
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/product", h.createProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/products", h.GetProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", h.GetProductHandler).Methods(http.MethodGet)
}

// handler for createProduct
func (h *Handler) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload

	// Parse the request body into CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Validate the payload using the Validate method
	if err := payload.Validate(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Convert payload to a Product object
	product := types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		CreatedAt:   time.Now(),
	}

	// Call the store method to create the product
	if err := h.store.CreateProduct(&product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}

// GetProductsHandler handles retrieving all products
func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with products
	utils.WriteJSON(w, http.StatusOK, products)
}

// GetProductHandler handles retrieving a product by its ID
func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Convert string to int
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Retrieve the product by ID
	product, err := h.store.GetProductById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// If product is not found, return a 404 error
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, errors.New("Product not found"))
		return
	}

	// Respond with the product
	utils.WriteJSON(w, http.StatusOK, product)
}
