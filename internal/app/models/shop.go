package models

import "github.com/jinzhu/gorm"

type Shop struct {
	gorm.Model
	Name  string `json:"name"`
	PAN   uint   `json:"pan"`
	Phone uint   `json:"phone"`
	Email string `json:"email"`
}

type ShopInterface interface {
	CreateShop(db *gorm.DB, data Shop) (*Shop, error)
	GetShop(db *gorm.DB) (*Shop, error)
}
