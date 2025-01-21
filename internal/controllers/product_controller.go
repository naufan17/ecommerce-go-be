package controllers

import (
	"ecommerce-go-be/internal/models"
	"ecommerce-go-be/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Products not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProduct(id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		} else if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.CreateProduct(product)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.UpdateProduct(product, id)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product"})
			return
		} else if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteProduct(id); err != nil {
		if err.Error() == "bad request" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
			return
		} else if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
