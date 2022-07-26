package utils

import (
	"billing-backend/internal/app/models"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

func GetExcelProducts(file multipart.File) (*[]models.ProductDetails, error) {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println("errors", err)
	}
	sheet1Name := "Sheet One"
	a := xlsx.GetSheetName(1)
	b := xlsx.GetRows(sheet1Name)
	fmt.Println("a", a)
	fmt.Println("b", b)
	rows := make([]map[string]interface{}, 0)
	details := make([]models.ProductDetails, 0)
	for i := 2; i < 5; i++ {
		row := map[string]interface{}{
			"Name":      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"Category":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"CostPrice": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			"SellPrice": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			"Quantity":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
		}
		rows = append(rows, row)
		costPrice, _ := strconv.ParseUint(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)), 10, 32)
		sellPrice, _ := strconv.ParseUint(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)), 10, 32)
		quantity, _ := strconv.ParseUint(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)), 10, 32)
		detail := models.ProductDetails{
			Name:      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			Category:  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			CostPrice: uint(costPrice),
			SellPrice: uint(sellPrice),
			Quantity:  uint(quantity),
		}
		details = append(details, detail)
	}

	fmt.Printf("%v \n", rows)
	fmt.Printf("%v \n", details)
	return &details, nil

}

func StoreToFile(fileDetails multipart.File) {
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(fileDetails)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
}

func GenerateBarCode(product models.ProductDetails) {

	fmt.Println("Generating code128 barcode for : ", product.Code)
	bcodes, err := code128.Encode(product.Code)

	if err != nil {
		fmt.Printf("String %s cannot be encoded", product.Code)
		os.Exit(1)
	}
	//barcode.Scale(bcodes, 250, 40)
	// scale to 250x20
	bcod, err := barcode.Scale(bcodes, 250, 50)
	fmt.Println(bcod)
	if err != nil {
		fmt.Println("Code128 scaling error!")
		os.Exit(1)
	}

	// now we want to append the code at the bottom
	// of the Codabar

	// Create an new image with text data
	// From https://github.com/llgcode/draw2d.samples/tree/master/helloworld
	// Set the global folder for searching fonts
	draw2d.SetFontFolder(".")

	// Initialize the graphic context on an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 250, 50))

	// set background to white
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(img)

	gc.FillStroke()

	// Set the font Montereymbi.ttf
	gc.SetFontData(draw2d.FontData{"Monterey", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	// Set the fill text color to black
	gc.SetFillColor(image.Black)

	gc.SetFontSize(14)
	gc.FillStringAt(product.Name, 50, 20)

	// create a new blank image with white background
	newImg := imaging.New(300, 150, color.NRGBA{255, 255, 255, 255})

	//paste the codabar to new blank image
	newImg = imaging.Paste(newImg, bcod, image.Pt(50, 50))

	//paste the text to the new blank image
	newImg = imaging.Paste(newImg, img, image.Pt(50, 100))

	err = draw2dimg.SaveToPngFile("./code129.png", newImg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// everything ok
	fmt.Println("Code128 code generated and saved to code128.png")
}
func RandInt() string {
	rand.Seed(time.Now().UnixNano())
	min, max := 100000000, 999999999
	return strconv.Itoa(min + rand.Intn(max-min))
	//return int64(min + rand.Intn(max-min))
}
