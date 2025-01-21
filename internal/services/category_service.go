package services

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"errors"
	"strconv"
)

func GetCategories() ([]dtos.CategoryDTO, error) {
	categories, err := repositories.GetCategories()

	if err != nil {
		return nil, errors.New("not found")
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
		return dtos.CategoryDTO{}, errors.New("bad request")
	}

	category, err := repositories.GetCategoryByID(intID)

	if err != nil {
		return dtos.CategoryDTO{}, errors.New("not found")
	}

	return dtos.ToCategoryDTO(category), nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	if _, err := repositories.CreateCategory(category); err != nil {
		return category, errors.New("internal server error")
	}

	return category, nil
}

func UpdateCategory(category models.Category, id string) (models.Category, error) {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return category, errors.New("bad request")
	}

	if _, err := repositories.UpdateCategory(category, intID); err != nil {
		return category, errors.New("internal server error")
	}

	return category, nil
}

func DeleteCategory(id string) error {
	intID, err := strconv.Atoi(id)

	if err != nil {
		return errors.New("bad request")
	}

	if err := repositories.DeleteCategory(intID); err != nil {
		return errors.New("not found")
	}

	return nil
}
