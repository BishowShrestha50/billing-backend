package controllers

import (
	"billing-backend/internal/app/adapter"
	"billing-backend/internal/app/application/utils/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB *gorm.DB
}

// Router is routing settings
func Router() *gin.Engine {
	controller := Controller{}
	db, err := adapter.Initialize()
	if err != nil {
		log.Panic("unable to pg connection :: ", err)
	}
	controller.DB = db
	return GetRoutes(controller)
}
func GetRoutes(controller Controller) *gin.Engine {
	r := gin.Default()
	//r.POST("/monitoring/token", controller.GetToken)
	authRoutes := r.Group("/api").Use(middleware.AuthMiddleware())
	authRoutes.POST("/billing/products", controller.CreateProductDetails)
	authRoutes.GET("/billing/products", controller.GetFilterProducts)
	authRoutes.POST("/billing/products-excel", controller.CreateProductsExcel)
	authRoutes.POST("/billing/barcode", controller.GenerateBarCode)
	authRoutes.POST("/billing/barcodes", controller.GenerateMultipleBarCode)
	authRoutes.POST("/billing/invoice", controller.CreateInvoice)
	authRoutes.GET("/billing/invoice", controller.GetAllInvoice)
	authRoutes.POST("/billing/generate-invoice/:id", controller.GenerateInvoice)
	// authRoutes.POST("/billing/shop", controller.CreateShop)
	// authRoutes.POST("/billing/shop", controller.GetShop)
	return r
}
