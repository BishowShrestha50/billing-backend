package controllers

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Details struct {
	Name      string
	Category  string
	CostPrice string
	SellPrice string
	Code      string
}

type M map[string]interface{}

var data = []map[string]interface{}{
	M{"Name": "Noval", "Category": "jeans", "CostPrice": "1200", "SellPrice": "1500", "Code": "123456"},
	M{"Name": "Yasa", "Category": "shirt", "CostPrice": "1000", "SellPrice": "1500", "Code": "678901"},
	M{"Name": "Nesa", "Category": "tshirt", "CostPrice": "1200", "SellPrice": "1800", "Code": "342122"},
}

func Excel() {
	// xlsx := excelize.NewFile()

	// sheet1Name := "Sheet One"
	// xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	// xlsx.SetCellValue(sheet1Name, "A1", "Name")
	// xlsx.SetCellValue(sheet1Name, "B1", "Category")
	// xlsx.SetCellValue(sheet1Name, "C1", "CostPrice")
	// xlsx.SetCellValue(sheet1Name, "D1", "SellPrice")
	// xlsx.SetCellValue(sheet1Name, "E1", "Code")
	// err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	// if err != nil {
	// 	log.Fatal("ERROR", err.Error())
	// }
	// for i, each := range data {
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Category"])
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["CostPrice"])
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), each["SellPrice"])
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each["Code"])
	// }
	// err = xlsx.SaveAs("./file2.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	xlsx, err := excelize.OpenFile("./file2.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	sheet1Name := "Sheet One"
	a := xlsx.GetSheetName(2)
	b := xlsx.GetRows(sheet1Name)
	fmt.Println("a", a)
	fmt.Println("b", b)
	rows := make([]M, 0)
	details := make([]Details, 0)
	for i := 2; i < 5; i++ {
		row := M{
			"Name":      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"Category":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"CostPrice": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			"SellPrice": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			"Code":      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
		}
		rows = append(rows, row)
		detail := Details{
			Name:      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			Category:  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			CostPrice: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			SellPrice: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			Code:      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
		}
		details = append(details, detail)
	}

	fmt.Printf("%v \n", rows)
	fmt.Printf("%v \n", details)
	byteData, err := json.Marshal(details[0])
	if err != nil {

	}
	productDetails := string(byteData)
	fmt.Printf("%v \n", productDetails)
	GenerateBarCode(details[0])
}

func GenerateBarCode(details Details) {

	code := details.Code
	fmt.Println("Generating code128 barcode for : ", code)
	bcodes, err := code128.Encode(code)

	if err != nil {
		fmt.Printf("String %s cannot be encoded", code)
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
	gc.FillStringAt(details.Name, 50, 20)

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
