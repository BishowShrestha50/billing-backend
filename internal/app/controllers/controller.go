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
	authRoutes.POST("/monitoring/insights/:id", controller.GetEnvironmentInsights)
	authRoutes.GET("/monitoring/podlist/:id", controller.GetPodList)
	authRoutes.GET("/monitoring/containerlist/:id", controller.GetContainerList)
	//authRoutes.GET("/monitoring/containermonitor/:id", controller.GetContainerMonitor)
	authRoutes.GET("/monitoring/events/:id", controller.GetEvents)
	authRoutes.POST("/monitoring/terminal", controller.GenerateToken)
	return r
}
