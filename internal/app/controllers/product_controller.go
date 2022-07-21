package controllers

import (
	"billing-backend/internal/app/application/utils"
	"billing-backend/internal/app/models"
	"fmt"
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

func (ctrl Controller) CreateProductsExcel(c *gin.Context) {

	file, _ := c.FormFile("file")
	fileDetails, _ := file.Open()
	defer fileDetails.Close()
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)
	utils.StoreToFile(fileDetails)
	products, _ := utils.GetExcelProducts("")
	fmt.Println(products)
	//   ctrl.Repo.CreateProductDetails(input)
	// return that we have successfully uploaded our file!
	c.JSON(http.StatusOK, "Successfully Uploaded File\n")
}

func (ctrl Controller) GenerateBarCode(c *gin.Context) {
	code := c.Param("code")
	name := c.Param("name")
	utils.GenerateBarCode(code, name)
}
