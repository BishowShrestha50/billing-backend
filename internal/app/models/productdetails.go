package models

import (
	"github.com/jinzhu/gorm"
)

type ProductDetails struct {
	gorm.Model
	Name      string `json:"name"`
	Category  string `json:"category"`
	CostPrice uint   `json:"cost_price"`
	SellPrice uint   `json:"sell_price"`
	Quantity  uint   `json:"quantity"`
	Code      string `json:"code"`
}
type ProductDetailsInterface interface {
	CreateProductDetails(db *gorm.DB, data ProductDetails) (*ProductDetails, error)
	GetAllProduct(db *gorm.DB) (*[]ProductDetails, error)
	FilterAllProduct(db *gorm.DB, page, perPage uint64, q, sortColumn, sortDirection string) (*map[string]interface{}, error)
}
