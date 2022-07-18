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
func (r *Repository) GetByID(eid uint64) (*models.ProductDetails, error) {
	data := models.ProductDetails{}

	return &data, nil
}
