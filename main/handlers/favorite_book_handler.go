package handlers

import (
	"errors"
	"net/http"

	"go-learn/main/models"
	"go-learn/main/requests"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFavoriteBooks(c *gin.Context) {
	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var favorites []models.FavoriteBook
	if err := preloadFavoriteBook(db).Where("user_id = ?", claims.UserID).Find(&favorites).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewFavoriteBookResponses(favorites))
}

func AddFavoriteBook(c *gin.Context) {
	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	var req requests.FavoriteBookCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	favorite := models.FavoriteBook{
		UserID: claims.UserID,
		BookID: req.BookID,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := ensureBookExists(tx, req.BookID); err != nil {
			return err
		}

		if err := tx.Create(&favorite).Error; err != nil {
			if isUniqueConstraintError(err) {
				return errors.New("book already in favorites")
			}

			return err
		}

		return preloadFavoriteBook(tx).First(&favorite, favorite.ID).Error
	}); err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusCreated, requests.NewFavoriteBookResponse(favorite))
}

func DeleteFavoriteBook(c *gin.Context) {
	claims, ok := currentAuthClaims(c)
	if !ok {
		return
	}

	bookID, ok := parseUintParam(c, "bookId")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var favorite models.FavoriteBook
	if err := preloadFavoriteBook(db).Where("user_id = ? AND book_id = ?", claims.UserID, bookID).First(&favorite).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	if err := db.Delete(&favorite).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewFavoriteBookResponse(favorite))
}
