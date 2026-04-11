package main

import (
	"go-learn/main/database"
	"go-learn/main/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Printf("diary database is unavailable: %v", err)
	}

	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.Run(":8080")
}
