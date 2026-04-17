package handlers

import (
	"errors"
	"net/http"
	"strings"

	"go-learn/main/models"
	"go-learn/main/requests"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	db := appDB(c)
	if db == nil {
		return
	}

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewCategoryResponses(categories))
}

func CreateCategory(c *gin.Context) {
	var req requests.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	category := models.Category{Name: name}
	if err := db.Create(&category).Error; err != nil {
		if isUniqueConstraintError(err) {
			writeStoreError(c, errors.New("category already exists"))
			return
		}

		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusCreated, requests.NewCategoryResponse(category))
}

func GetCategoryByID(c *gin.Context) {
	categoryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var category models.Category
	if err := db.First(&category, categoryID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewCategoryResponse(category))
}

func UpdateCategory(c *gin.Context) {
	categoryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req requests.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var category models.Category
	if err := db.First(&category, categoryID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	category.Name = name
	if err := db.Save(&category).Error; err != nil {
		if isUniqueConstraintError(err) {
			writeStoreError(c, errors.New("category already exists"))
			return
		}

		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewCategoryResponse(category))
}

func DeleteCategory(c *gin.Context) {
	categoryID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var category models.Category
	if err := db.First(&category, categoryID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewCategoryResponse(category))
}
