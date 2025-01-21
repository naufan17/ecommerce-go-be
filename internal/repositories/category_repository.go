package repositories

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/utils"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category

	if err := utils.DB.Select("id", "name").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryByID(id int) (models.Category, error) {
	var category models.Category

	if err := utils.DB.Select("id", "name").First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	if err := utils.DB.Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategory(category models.Category, id int) (models.Category, error) {
	if err := utils.DB.Where("id = ?", id).Updates(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func DeleteCategory(id int) error {
	var category models.Category

	if err := utils.DB.Delete(&category, id).Error; err != nil {
		return err
	}

	return nil
}
