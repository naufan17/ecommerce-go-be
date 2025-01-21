package dtos

import (
	"ecommerce-go-be/internal/models"
)

type ProductDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
}

func ToProductDTO(product models.Product) ProductDTO {
	return ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Category:    product.Category.Name,
	}
}
