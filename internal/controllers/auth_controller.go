package controllers

import (
	"ecommerce-go-be/internal/dtos"
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.Register(user)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user"})
			return
		} else if err.Error() == "conflict" {
			c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user dtos.LoginDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := services.Login(user)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accessToken})
}
