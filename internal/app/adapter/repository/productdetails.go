package repository

import (
	"billing-backend/internal/app/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) models.ProductDetailsInterface {
	return &Repository{DB: db}
}
func (r *Repository) CreateProductDetails(data models.ProductDetails) (*models.ProductDetails, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
