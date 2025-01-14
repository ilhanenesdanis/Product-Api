package main

import (
	"fmt"
	"log"
	"product-api/config"
	"product-api/internal/domain"
)

func main() {
	db, err := config.GenerateDb()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("Running migrations...")

	err = db.AutoMigrate(&domain.Product{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migration completed")
}
