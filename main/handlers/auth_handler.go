package handlers

import (
	"go-learn/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	user, err := registerUser(db, req.Username, req.Password, req.Role)
	if err != nil {
		writeAuthError(c, err)
		return
	}

	token, err := generateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, registerResponse{
		Token: token,
		User:  newAuthUserDTO(*user),
	})
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	db := appDB(c)
	if db == nil {
		return
	}

	user, err := authenticateUser(db, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := generateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

func newAuthUserDTO(user models.User) authUserDTO {
	return authUserDTO{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
