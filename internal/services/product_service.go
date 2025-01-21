package services

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"errors"
	"strconv"
)

func GetProducts() ([]dtos.ProductDTO, error) {
	products, err := repositories.GetProducts()

	if err != nil {
		return nil, errors.New("not found")
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
		return dtos.ProductDTO{}, errors.New("bad request")
	}

	product, err := repositories.GetProductByID(intID)

	if err != nil {
		return dtos.ProductDTO{}, errors.New("not found")
	}

	return dtos.ToProductDTO(product), nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	if _, err := repositories.CreateProduct(product); err != nil {
		return product, errors.New("internal server error")
	}

	return product, nil
}

func UpdateProduct(product models.Product, id string) (models.Product, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return product, errors.New("bad request")
	}

	if _, err := repositories.UpdateProduct(product, intID); err != nil {
		return product, errors.New("internal server error")
	}

	return product, nil
}

func DeleteProduct(id string) error {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return errors.New("bad request")
	}

	if err := repositories.DeleteProduct(intID); err != nil {
		return errors.New("not found")
	}

	return nil
}
