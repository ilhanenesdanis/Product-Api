package http

import (
	"product-api/internal/usecase"

	"github.com/gorilla/mux"
)

func NewRouter(productUsecase usecase.ProductUsecase) *mux.Router {
	router := mux.NewRouter()

	productHandler := NewProductHandler(productUsecase)

	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct).Methods("DELETE")

	return router
}
