package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string  `json:"name"`
	NoHandphone string  `json:"no_handphone"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Address     string  `json:"address"`
	Order       []Order `gorm:"foreignKey:UserID;references:ID"`
}
