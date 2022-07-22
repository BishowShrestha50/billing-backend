package controllers

import (
	"billing-backend/internal/app/adapter"
	"billing-backend/internal/app/adapter/repository"
	"billing-backend/internal/app/application/utils/middleware"
	"billing-backend/internal/app/models"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	base = adapter.BaseRepo{}
)

type Controller struct {
	Repo models.ProductDetailsInterface
}

// Router is routing settings
func Router() *gin.Engine {
	controller := Controller{}
	err := base.Initialize()
	db := repository.NewRepo(base.DB)
	controller.Repo = db
	if err != nil {
		log.Panic("unable to pg connection :: ", err)
	}
	return GetRoutes(controller)
}
func GetRoutes(controller Controller) *gin.Engine {
	r := gin.Default()
	//r.POST("/monitoring/token", controller.GetToken)
	authRoutes := r.Group("/").Use(middleware.AuthMiddleware(controller.Repo))
	authRoutes.GET("/billing/products", controller.CreateProductDetails)

	return r
}
