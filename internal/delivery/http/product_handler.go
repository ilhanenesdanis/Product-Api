package http

import (
	"encoding/json"
	"net/http"
	"product-api/internal/domain"
	"product-api/internal/usecase"
	"strconv"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

// tüm ürünleri getirir
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.ProductUsecase.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}
	product, err := ph.ProductUsecase.GetProductByID(conv)
	if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := ph.ProductUsecase.CreateProduct(&product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product.ID = int(conv)
	if err := h.ProductUsecase.UpdateProduct(&product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.ProductUsecase.DeleteProduct(conv); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
