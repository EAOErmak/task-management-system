package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-learn/main/database"
	"go-learn/main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func appDB(c *gin.Context) *gorm.DB {
	if database.DB == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "database is not initialized"})
		return nil
	}

	return database.DB
}

func parseUintParam(c *gin.Context, name string) (uint, bool) {
	raw := strings.TrimSpace(c.Param(name))
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid %s", name)})
		return 0, false
	}

	return uint(id), true
}

func currentAuthClaims(c *gin.Context) (*authClaims, bool) {
	rawClaims, exists := c.Get(authClaimsContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth context"})
		return nil, false
	}

	claims, ok := rawClaims.(*authClaims)
	if !ok || claims == nil || claims.UserID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth context"})
		return nil, false
	}

	return claims, true
}

func preloadBook(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Author").
		Preload("Category")
}

func preloadFavoriteBook(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Book.Author").
		Preload("Book.Category")
}

func ensureAuthorExists(tx *gorm.DB, authorID uint) error {
	return ensureRecordExists(tx, &models.Author{}, authorID, "author")
}

func ensureCategoryExists(tx *gorm.DB, categoryID uint) error {
	return ensureRecordExists(tx, &models.Category{}, categoryID, "category")
}

func ensureBookExists(tx *gorm.DB, bookID uint) error {
	return ensureRecordExists(tx, &models.Book{}, bookID, "book")
}

func ensureRecordExists(tx *gorm.DB, model any, id uint, entityName string) error {
	var count int64
	if err := tx.Model(model).Where("id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("%s not found", entityName)
	}

	return nil
}

func writeStoreError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
	case isValidationError(err):
		status := http.StatusBadRequest
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			status = http.StatusConflict
		}

		c.JSON(status, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func isValidationError(err error) bool {
	if err == nil {
		return false
	}

	message := strings.ToLower(err.Error())

	return strings.Contains(message, "required") ||
		strings.Contains(message, "invalid") ||
		strings.Contains(message, "cannot") ||
		strings.Contains(message, "already exists") ||
		strings.Contains(message, "must be") ||
		strings.Contains(message, "not found")
}

func isUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	message := strings.ToLower(err.Error())

	return strings.Contains(message, "duplicate key") ||
		strings.Contains(message, "unique constraint") ||
		strings.Contains(message, "violates unique")
}
