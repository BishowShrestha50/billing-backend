package models

type ProductDetails struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	CostPrice string `json:"cost_price"`
	SellPrice string `json:"sell_price"`
	Quantity  string `json:"quantity"`
	Code      string `json:"code"`
}
type ProductDetailsInterface interface {
	CreateProductDetails(data ProductDetails) (*ProductDetails, error)
}
