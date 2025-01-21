package services

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"strconv"
)

func GetProducts() ([]dtos.ProductDTO, error) {
	products, err := repositories.GetProducts()

	if err != nil {
		return nil, err
	}

	var productDTOs []dtos.ProductDTO

	for _, product := range products {
		productDTOs = append(productDTOs, dtos.ToProductDTO(product))
	}

	return productDTOs, nil
}

func GetProduct(id string) (dtos.ProductDTO, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return dtos.ProductDTO{}, err
	}

	product, err := repositories.GetProductByID(intID)

	if err != nil {
		return dtos.ProductDTO{}, err
	}

	return dtos.ToProductDTO(product), nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	createdProduct, err := repositories.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}

func UpdateProduct(product models.Product, id string) (models.Product, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return models.Product{}, err
	}

	updatedProduct, err := repositories.UpdateProduct(product, intID)

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

	if err := repositories.DeleteProduct(intID); err != nil {
		return err
	}

	return nil
}
