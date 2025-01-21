package middlewares

import (
	"ecommerce-go-be/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateJWT(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		next(c)
	}
}
