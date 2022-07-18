package models

import (
	"time"
)

type ProductDetails struct {
	Name      string      `json:"name"`
	Status    interface{} `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}
type ProductDetailsInterface interface {
	GetByID(eid uint64) (*ProductDetails, error)
}
