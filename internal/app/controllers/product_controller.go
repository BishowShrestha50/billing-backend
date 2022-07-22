package controllers

import (
	"billing-backend/internal/app/application/utils"
	"billing-backend/internal/app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ctrl Controller) CreateProductDetails(c *gin.Context) {

	var input models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Code = utils.RandInt()
	data, err := ctrl.Repo.CreateProductDetails(input)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}

func (ctrl Controller) CreateProductsExcel(c *gin.Context) {

	file, _ := c.FormFile("file")
	fileDetails, _ := file.Open()
	defer fileDetails.Close()
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)
	// utils.StoreToFile(fileDetails)
	products, _ := utils.GetExcelProducts(fileDetails)
	fmt.Println(products)
	for _, item := range *products {
		_, err := ctrl.Repo.CreateProductDetails(item)
		if err != nil {
			continue
		}
	}
	c.JSON(http.StatusOK, "Successfully Uploaded File\n")
}

func (ctrl Controller) GenerateBarCode(c *gin.Context) {
	var input models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.GenerateBarCode(input)
}
func (ctrl Controller) GenerateMultipleBarCode(c *gin.Context) {
	var input []models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, barcode := range input {
		barcode.Code = utils.RandInt()
		utils.GenerateBarCode(barcode)
	}

}
