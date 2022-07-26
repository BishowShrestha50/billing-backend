package models

import "github.com/jinzhu/gorm"

type Invoice struct {
	gorm.Model
	Total        float64        `json:"total"`
	SubTotal     float64        `json:"sub_total"`
	Tax          uint           `json:"tax"`
	Discount     uint           `json:"discount"`
	InvoiceItems []InvoiceItems `gorm:"foreignkey:InvoiceID" json:"invoice_items"`
}
type InvoiceItems struct {
	gorm.Model
	Invoice   *Invoice `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID uint     `gorm:"not null" json:"invoice_id"`
	Name      string   `json:"name"`
	Quantity  uint     `json:"quantity"`
	Price     float64  `json:"price"`
	Total     float64  `json:"total"`
}
type InvoiceInterface interface {
	CreateInvoice(db *gorm.DB, data Invoice) (*Invoice, error)
	GetAllInvoice(db *gorm.DB) (*[]Invoice, error)
	GetByInvoiceInvoiceItems(db *gorm.DB, id uint) (*[]InvoiceItems, error)
}
