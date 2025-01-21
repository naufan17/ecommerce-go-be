package repositories

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/utils"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product

	if err := utils.DB.Preload("Category").Select("id", "name", "description", "price", "quantity", "category_id").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product

	if err := utils.DB.Preload("Category").Select("id", "name", "description", "price", "quantity", "category_id").First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	if err := utils.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func UpdateProduct(product models.Product, id int) (models.Product, error) {
	if err := utils.DB.Where("id = ?", id).Updates(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func DeleteProduct(id int) error {
	var product models.Product

	if err := utils.DB.Delete(&product, id).Error; err != nil {
		return err
	}

	return nil
}
