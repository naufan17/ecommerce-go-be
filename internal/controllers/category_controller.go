package controllers

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Categories not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategory(id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
			return
		} else if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.CreateCategory(category)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update category"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.UpdateCategory(category, id)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update category"})
			return
		} else if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteCategory(id); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
			return
		} else if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
