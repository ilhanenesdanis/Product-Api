package main

import (
	"fmt"
	"product-api/config"
)

func main() {
	_, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Connection Error")
	}

}
