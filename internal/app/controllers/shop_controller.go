package controllers

import (
	"billing-backend/internal/app/adapter/repository"
	"billing-backend/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var shopInterface = repository.Newshop()

func (ctrl Controller) CreateShop(c *gin.Context) {
	var shop models.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := shopInterface.CreateShop(ctrl.DB, shop)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, data)
}
func (ctrl Controller) GetShop(c *gin.Context) {
	data, err := shopInterface.GetShop(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, data)
}
