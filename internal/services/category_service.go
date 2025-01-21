package services

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"strconv"
)

func GetCategories() ([]dtos.CategoryDTO, error) {
	categories, err := repositories.GetCategories()

	if err != nil {
		return nil, err
	}

	var categoryDTOs []dtos.CategoryDTO

	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, dtos.ToCategoryDTO(category))
	}

	return categoryDTOs, nil
}

func GetCategory(id string) (dtos.CategoryDTO, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return dtos.CategoryDTO{}, err
	}

	category, err := repositories.GetCategoryByID(intID)

	if err != nil {
		return dtos.CategoryDTO{}, err
	}

	return dtos.ToCategoryDTO(category), nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	createdCategory, err := repositories.CreateCategory(category)

	if err != nil {
		return models.Category{}, err
	}

	return createdCategory, nil
}

func UpdateCategory(category models.Category, id string) (models.Category, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return models.Category{}, err
	}

	updatedCategory, err := repositories.UpdateCategory(category, intID)

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

	if err := repositories.DeleteCategory(intID); err != nil {
		return err
	}

	return nil
}
