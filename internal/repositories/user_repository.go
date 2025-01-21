package repositories

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/utils"
)

func CreateUser(user models.User) (models.User, error) {
	if err := utils.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if err := utils.DB.Where("email = ?", email).Select("email", "password").First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
