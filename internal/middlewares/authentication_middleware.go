package middlewares

import (
	"ecommerce-go-be/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
		c.Abort()
		return
	}

	if _, err := utils.ValidateJWT(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
