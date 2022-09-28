package adapter

import (
	"billing-backend/internal/app/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Initialize() (*gorm.DB, error) {
	url := fmt.Sprintf("host=%v port=%s user=%v password=%v dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	//url := os.Getenv("DB_SOURCE")
	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.ProductDetails{}, &models.Invoice{}, &models.InvoiceItems{})
	return db, nil
}
