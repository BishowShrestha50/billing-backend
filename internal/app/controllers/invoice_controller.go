package controllers

import (
	"billing-backend/internal/app/adapter/repository"
	"billing-backend/internal/app/models"
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var invoiceInterface = repository.NewInvoice()

func (ctrl Controller) CreateInvoice(c *gin.Context) {

	var input models.Invoice
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("hhh", input)
	data, err := invoiceInterface.CreateInvoice(ctrl.DB, input)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}

func (ctrl Controller) GetAllInvoice(c *gin.Context) {
	data, err := invoiceInterface.GetAllInvoice(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}
func (ctrl Controller) GetInvoiceItems(c *gin.Context) {
	id := c.Param("id")
	invoiceID, _ := strconv.ParseUint(id, 10, 32)
	data, err := invoiceInterface.GetByInvoiceInvoiceItems(ctrl.DB, uint(invoiceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, "pass")
}

func (ctrl Controller) GenerateInvoice(c *gin.Context) {
	fmt.Println("sdfdf")
	id := c.Param("id")
	fmt.Println("sss", id)
	invoiceID, _ := strconv.ParseUint(id, 10, 32)
	data, err := invoiceInterface.GetInvoiceByID(ctrl.DB, uint(invoiceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	templateData := struct {
		InvoiceID    uint
		CreatedDate  string
		Total        float64
		SubTotal     float64
		Tax          uint
		Discount     uint
		InvoiceItems []models.InvoiceItems
	}{
		InvoiceID:    data.ID,
		CreatedDate:  data.CreatedAt.String(),
		Total:        data.SubTotal,
		SubTotal:     data.SubTotal,
		Tax:          data.Tax,
		Discount:     data.Discount,
		InvoiceItems: data.InvoiceItems,
	}
	t, err := template.ParseFiles("./design/index.html")
	if err != nil {
		logrus.Error("tmp-err", err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		logrus.Error("err", err)
	}
	//r.body = buf.String()
	f, err := os.Create("./output/index.html")
	if err != nil {
		logrus.Error("eee", err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(buf.String()))
	w.Flush()
}
