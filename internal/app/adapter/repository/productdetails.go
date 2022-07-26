package repository

import (
	"billing-backend/internal/app/models"
	"math"

	"github.com/jinzhu/gorm"
)

type Products struct {
}

func NewProduct() models.ProductDetailsInterface {
	return &Products{}
}
func (r *Products) CreateProductDetails(db *gorm.DB, data models.ProductDetails) (*models.ProductDetails, error) {
	err := db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Products) GetAllProduct(db *gorm.DB) (*[]models.ProductDetails, error) {
	data := []models.ProductDetails{}
	err := db.Model(&models.ProductDetails{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Products) FilterAllProduct(db *gorm.DB, page, perPage uint64, q, sortColumn, sortDirection string) (*map[string]interface{}, error) {
	products := []models.ProductDetails{}
	var total int64
	f := db.Model(&models.ProductDetails{})
	f = f.Where("lower(name) LIKE lower(?)", "%"+q+"%")
	err := f.Order(sortColumn + " " + sortDirection).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	f.Count(&total)
	return &map[string]interface{}{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / int64(perPage))),
	}, nil

}
