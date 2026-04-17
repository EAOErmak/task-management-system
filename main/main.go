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

	api := router.Group("/")
	api.Use(handlers.AuthMiddleware())
	{
		api.GET("/authors", handlers.GetAllAuthors)
		api.POST("/authors", handlers.CreateAuthor)
		api.GET("/authors/:id", handlers.GetAuthorByID)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", handlers.DeleteAuthor)

		api.GET("/categories", handlers.GetAllCategories)
		api.POST("/categories", handlers.CreateCategory)
		api.GET("/categories/:id", handlers.GetCategoryByID)
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)

		api.GET("/books", handlers.GetAllBooks)
		api.POST("/books", handlers.CreateBook)
		api.GET("/books/:id", handlers.GetBookByID)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)

		api.GET("/favorite-books", handlers.GetFavoriteBooks)
		api.POST("/favorite-books", handlers.AddFavoriteBook)
		api.DELETE("/favorite-books/:bookId", handlers.DeleteFavoriteBook)
	}

	router.Run(":8080")
}
