package handlers

import (
	"net/http"
	"strings"

	"go-learn/main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBooks(c *gin.Context) {
	db := appDB(c)
	if db == nil {
		return
	}

	var books []models.Book
	if err := preloadBook(db).Find(&books).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var req bookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	title := strings.TrimSpace(req.Title)
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if req.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than or equal to 0"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	book := models.Book{
		Title:      title,
		AuthorID:   req.AuthorID,
		CategoryID: req.CategoryID,
		Price:      req.Price,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := ensureAuthorExists(tx, req.AuthorID); err != nil {
			return err
		}

		if err := ensureCategoryExists(tx, req.CategoryID); err != nil {
			return err
		}

		if err := tx.Create(&book).Error; err != nil {
			return err
		}

		return preloadBook(tx).First(&book, book.ID).Error
	}); err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	bookID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var book models.Book
	if err := preloadBook(db).First(&book, bookID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	bookID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req bookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	title := strings.TrimSpace(req.Title)
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if req.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than or equal to 0"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var book models.Book
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&book, bookID).Error; err != nil {
			return err
		}

		if err := ensureAuthorExists(tx, req.AuthorID); err != nil {
			return err
		}

		if err := ensureCategoryExists(tx, req.CategoryID); err != nil {
			return err
		}

		book.Title = title
		book.AuthorID = req.AuthorID
		book.CategoryID = req.CategoryID
		book.Price = req.Price

		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		return preloadBook(tx).First(&book, book.ID).Error
	}); err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	bookID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var book models.Book
	if err := preloadBook(db).First(&book, bookID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}
