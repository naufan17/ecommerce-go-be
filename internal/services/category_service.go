package services

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"strconv"
)

func GetCategories() ([]models.Category, error) {
	return repositories.GetCategories()
}

func GetCategory(id string) (models.Category, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return models.Category{}, err
	}

	return repositories.GetCategoryByID(intID)
}

func CreateCategory(category models.Category) (models.Category, error) {
	createdCategory, err := repositories.CreateCategory(category)

	if err != nil {
		return models.Category{}, err
	}

	return createdCategory, nil
}

func UpdateCategory(category models.Category) (models.Category, error) {
	updatedCategory, err := repositories.UpdateCategory(category)

	if err != nil {
		return models.Category{}, err
	}

	return updatedCategory, nil
}

func DeleteCategory(id string) error {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	return repositories.DeleteCategory(intID)
}
