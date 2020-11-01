package models

import (
	"gorm.io/gorm"
)

// User -> model for user entity
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
}
