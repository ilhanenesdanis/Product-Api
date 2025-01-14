package http

import (
	"encoding/json"
	"net/http"
	"product-api/internal/usecase"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

// getAllProducts tüm ürünleri getirir
func (ph *ProductHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.ProductUsecase.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
