package controllers

import (
	"billing-backend/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl Controller) GetProductDetails(c *gin.Context) {

	var input models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := ctrl.Repo.CreateProductDetails(input)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}
