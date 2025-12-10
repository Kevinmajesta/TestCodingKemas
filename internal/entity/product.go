package entity

import (
	"github.com/google/uuid"
)

type Products struct {
	ProductID uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"not null;unique"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null;check:price >= 0"`
	Stock     int       `json:"stock" gorm:"not null;default:0;check:stock >= 0"`
}

func NewProduct(name string, price float64, stock int) *Products {
	return &Products{
		ProductID: uuid.New(),
		Name:      name,
		Price:     price,
		Stock:     stock,
	}
}

func UpdateProduct(productID uuid.UUID, name string, price float64, stock int) *Products {
	return &Products{
		ProductID: productID,
		Name:      name,
		Price:     price,
		Stock:     stock,
	}
}
