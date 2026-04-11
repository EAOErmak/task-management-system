package models

import (
	"errors"
	"strings"
)

const DefaultUserRole = "user"

type User struct {
	BaseModel
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"-"`
	Role     string `gorm:"column:role;not null" json:"role"`
}

func (User) TableName() string {
	return "users"
}

func NewUser(username, passwordHash, role string) (*User, error) {
	trimmedUsername := strings.TrimSpace(username)
	if trimmedUsername == "" {
		return nil, errors.New("username is required")
	}

	if strings.TrimSpace(passwordHash) == "" {
		return nil, errors.New("password is required")
	}

	trimmedRole := strings.TrimSpace(role)
	if trimmedRole == "" {
		trimmedRole = DefaultUserRole
	}

	return &User{
		Username: trimmedUsername,
		Password: passwordHash,
		Role:     trimmedRole,
	}, nil
}
