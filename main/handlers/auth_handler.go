package handlers

import (
	"net/http"

	"go-learn/main/requests"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req requests.UserCreateRequest
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

	c.JSON(http.StatusCreated, requests.RegisterResponse{
		Token: token,
		User:  requests.NewUserResponse(*user),
	})
}

func Login(c *gin.Context) {
	var req requests.LoginRequest
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

	c.JSON(http.StatusOK, requests.LoginResponse{Token: token})
}
