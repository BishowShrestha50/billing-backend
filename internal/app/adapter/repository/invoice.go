package repository

import (
	"billing-backend/internal/app/models"

	"github.com/jinzhu/gorm"
)

type Invoice struct {
}

func NewInvoice() models.InvoiceInterface {
	return &Invoice{}
}
func (r *Invoice) CreateInvoice(db *gorm.DB, data models.Invoice) (*models.Invoice, error) {
	err := db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r *Invoice) GetAllInvoice(db *gorm.DB) (*[]models.Invoice, error) {
	data := []models.Invoice{}
	err := db.Model(&models.Invoice{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Invoice) GetByInvoiceInvoiceItems(db *gorm.DB, id uint) (*[]models.InvoiceItems, error) {
	data := []models.InvoiceItems{}
	err := db.Model(&models.InvoiceItems{}).Where("invoice_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Invoice) GetInvoiceByID(db *gorm.DB, id uint) (*models.Invoice, error) {
	data := models.Invoice{}
	err := db.Model(&models.Invoice{}).Preload("InvoiceItems").Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
