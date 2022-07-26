package controllers

import (
	"billing-backend/internal/app/adapter/repository"
	"billing-backend/internal/app/application/utils"
	"billing-backend/internal/app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var productInterface = repository.NewProduct()

func (ctrl Controller) CreateProductDetails(c *gin.Context) {

	var input models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Code = utils.RandInt()
	fmt.Println("hhh", input)
	data, err := productInterface.CreateProductDetails(ctrl.DB, input)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}

func (ctrl Controller) GetAllProduct(c *gin.Context) {
	data, err := productInterface.GetAllProduct(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}

func (ctrl Controller) CreateProductsExcel(c *gin.Context) {
	logrus.Println("hhghh")
	file, _ := c.FormFile("myFile")
	fileDetails, _ := file.Open()
	defer fileDetails.Close()
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)
	// utils.StoreToFile(fileDetails)
	products, _ := utils.GetExcelProducts(fileDetails)
	fmt.Println(products)
	for _, item := range *products {
		item.Code = utils.RandInt()
		_, err := productInterface.CreateProductDetails(ctrl.DB, item)
		if err != nil {
			continue
		}
	}
	c.JSON(http.StatusOK, "Successfully Uploaded File\n")
}

func (ctrl Controller) GetFilterProducts(c *gin.Context) {
	page, _ := strconv.ParseUint(c.Query("page"), 10, 32)
	size, _ := strconv.ParseUint(c.Query("size"), 10, 32)
	search := c.Query("search")
	sortColumn := c.Query("sort-column")
	sortDirection := c.Query("sort-direction")
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	if sortColumn == "" {
		sortColumn = "id"
	}
	if sortDirection == "" {
		sortDirection = "desc"
	}
	datas, err := productInterface.FilterAllProduct(ctrl.DB, page, size, search, sortColumn, sortDirection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, datas)
}

func (ctrl Controller) GenerateBarCode(c *gin.Context) {
	var input models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("asass", input)
	utils.GenerateBarCode(input)
}
func (ctrl Controller) GenerateMultipleBarCode(c *gin.Context) {
	var input []models.ProductDetails
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, barcode := range input {
		utils.GenerateBarCode(barcode)
	}

}
