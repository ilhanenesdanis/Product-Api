package main

import (
	"fmt"
	"log"
	nethttp "net/http"
	"product-api/config"
	"product-api/internal/delivery/http"
	"product-api/internal/repository"
	"product-api/internal/usecase"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Connection Error")
	}
	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)

	router := http.NewRouter(productUsecase)

	log.Println("Server started at :8080")
	log.Fatal(nethttp.ListenAndServe(":8080", router))
}
