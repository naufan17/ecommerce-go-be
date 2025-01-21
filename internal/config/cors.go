package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Origin, Content-Length, Content-Type, Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	})
}
