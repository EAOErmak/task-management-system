package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"go-learn/main/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	defaultJWTSecret = "secret"
	jwtLifetime      = time.Hour
)

type authClaims struct {
	Username string `json:"username"`
	Role     string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func registerUser(db *gorm.DB, username, password, role string) (*models.User, error) {
	trimmedUsername := strings.TrimSpace(username)
	if trimmedUsername == "" {
		return nil, errors.New("username is required")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}

	var count int64
	if err := db.Model(&models.User{}).Where("username = ?", trimmedUsername).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("username already exists")
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := models.NewUser(trimmedUsername, passwordHash, role)
	if err != nil {
		return nil, err
	}

	if err := db.Create(user).Error; err != nil {
		if isUniqueConstraintError(err) {
			return nil, errors.New("username already exists")
		}

		return nil, err
	}

	return user, nil
}

func authenticateUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", strings.TrimSpace(username)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func generateToken(user models.User) (string, error) {
	now := time.Now().UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtLifetime)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})

	return token.SignedString(jwtSecret())
}

func parseToken(tokenString string) (*authClaims, error) {
	claims := &authClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}

		return jwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid || strings.TrimSpace(claims.Username) == "" {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func jwtSecret() []byte {
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if secret == "" {
		secret = defaultJWTSecret
	}

	return []byte(secret)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func writeAuthError(c *gin.Context, err error) {
	switch {
	case err == nil:
		return
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
