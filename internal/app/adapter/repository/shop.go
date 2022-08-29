package repository

import (
	"billing-backend/internal/app/models"

	"github.com/jinzhu/gorm"
)

type Shop struct {
}

func Newshop() models.ShopInterface {
	return &Shop{}
}
func (r *Shop) CreateShop(db *gorm.DB, data models.Shop) (*models.Shop, error) {
	err := db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r *Shop) GetShop(db *gorm.DB) (*models.Shop, error) {
	data := models.Shop{}
	err := db.Model(&models.Shop{}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
