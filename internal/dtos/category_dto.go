package dtos

import (
	"ecommerce-go-be/internal/models"
)

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToCategoryDTO(category models.Category) CategoryDTO {
	return CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
