package main

import (
	"To-Do/internal/models"
	"To-Do/pkg/apiserver"
	"To-Do/pkg/handler"
	"fmt"
	"log"
	_ "To-Do/docs"
)


// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8000
// @BasePath /


func main() {
	err := models.Migrate()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		fmt.Println("Migration completed successfully.")
	}

	server := new(apiserver.ApiServer)
	handlers := new(handler.Handler)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}




