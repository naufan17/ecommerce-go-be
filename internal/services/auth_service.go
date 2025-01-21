package services

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/repositories"
	"ecommerce-go-be/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	if _, err := repositories.CreateUser(user); err != nil {
		return models.User{}, errors.New("conflict")
	}

	return user, nil
}

func Login(user dtos.LoginDTO) (dtos.AccessTokenDTO, error) {
	userFromDB, err := repositories.GetUserByEmail(user.Email)

	if err != nil {
		return dtos.AccessTokenDTO{}, errors.New("not found")
	}

	if !utils.ComparePassword(user.Password, userFromDB.Password) {
		return dtos.AccessTokenDTO{}, errors.New("unauthorized")
	}

	accessToken, expiresIn, tokenType, err := utils.GenerateJWT(userFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDTO{}, errors.New("internal server error")
	}

	return dtos.AccessTokenDTO{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
		TokenType:   tokenType,
	}, nil
}
