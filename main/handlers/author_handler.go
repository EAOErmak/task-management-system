package handlers

import (
	"errors"
	"net/http"
	"strings"

	"go-learn/main/models"
	"go-learn/main/requests"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {
	db := appDB(c)
	if db == nil {
		return
	}

	var authors []models.Author
	if err := db.Find(&authors).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewAuthorResponses(authors))
}

func CreateAuthor(c *gin.Context) {
	var req requests.AuthorCreateRequest
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

	author := models.Author{Name: name}
	if err := db.Create(&author).Error; err != nil {
		if isUniqueConstraintError(err) {
			writeStoreError(c, errors.New("author already exists"))
			return
		}

		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusCreated, requests.NewAuthorResponse(author))
}

func GetAuthorByID(c *gin.Context) {
	authorID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var author models.Author
	if err := db.First(&author, authorID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewAuthorResponse(author))
}

func UpdateAuthor(c *gin.Context) {
	authorID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req requests.AuthorCreateRequest
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

	var author models.Author
	if err := db.First(&author, authorID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	author.Name = name
	if err := db.Save(&author).Error; err != nil {
		if isUniqueConstraintError(err) {
			writeStoreError(c, errors.New("author already exists"))
			return
		}

		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewAuthorResponse(author))
}

func DeleteAuthor(c *gin.Context) {
	authorID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	var author models.Author
	if err := db.First(&author, authorID).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		writeStoreError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.NewAuthorResponse(author))
}
