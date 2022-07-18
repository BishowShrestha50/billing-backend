package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl Controller) GetProductDetails(c *gin.Context) {

	c.JSON(http.StatusOK, "pass")
}
