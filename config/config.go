package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "sqlserver://SA:Enes5858*@127.0.0.1:1433?database=ProductDB&trustservercertificate=true"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("connected to database")
	return db, nil
}

func GenerateDb() (*gorm.DB, error) {

	masterDSN := "sqlserver://SA:Enes5858*@127.0.0.1:1433?database=master&Encrypt=false&TrustServerCertificate=true"
	masterDB, err := gorm.Open(sqlserver.Open(masterDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to master database: %w", err)
	}

	targetDBName := "ProductDB"

	var dbExists int
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM sys.databases WHERE name = '%s'", targetDBName)
	if err := masterDB.Raw(checkQuery).Scan(&dbExists).Error; err != nil {
		return nil, fmt.Errorf("Failed to check database existence: %w", err)
	}

	if dbExists == 0 {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", targetDBName)
		if err := masterDB.Exec(createQuery).Error; err != nil {
			return nil, fmt.Errorf("Failed to create database: %w", err)
		}
		log.Printf("Database '%s' created successfully.", targetDBName)
	} else {
		log.Printf("Database '%s' already exists.", targetDBName)
	}

	targetDSN := fmt.Sprintf("sqlserver://SA:Enes5858*@127.0.0.1:1433?database=%s&Encrypt=false&TrustServerCertificate=true", targetDBName)
	targetDB, err := gorm.Open(sqlserver.Open(targetDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to target database: %w", err)
	}

	return targetDB, nil
}
