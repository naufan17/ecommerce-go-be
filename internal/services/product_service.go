package services

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"strconv"
)

func GetProducts() ([]models.Product, error) {
	return repositories.GetProducts()
}

func GetProduct(id string) (models.Product, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return models.Product{}, err
	}

	return repositories.GetProductByID(intID)
}

func CreateProduct(product models.Product) (models.Product, error) {
	createdProduct, err := repositories.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	updatedProduct, err := repositories.UpdateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func DeleteProduct(id string) error {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	return repositories.DeleteProduct(intID)
}
