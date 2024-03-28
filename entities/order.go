package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"foreignKey:ID"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
