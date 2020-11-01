package models

import (
	"github.com/jinzhu/gorm"
)

// Product -> model for Product entity
type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}
