package routes

import (
	"ecommerce-go-be/internal/controllers"
	"ecommerce-go-be/internal/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Origin, Content-Length, Content-Type, Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)

		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategory)
		api.POST("/categories", middlewares.AuthenticationMiddleware, controllers.CreateCategory)
		api.PUT("/categories/:id", middlewares.AuthenticationMiddleware, controllers.UpdateCategory)
		api.DELETE("/categories/:id", middlewares.AuthenticationMiddleware, controllers.DeleteCategory)

		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProduct)
		api.POST("/products", middlewares.AuthenticationMiddleware, controllers.CreateProduct)
		api.PUT("/products/:id", middlewares.AuthenticationMiddleware, controllers.UpdateProduct)
		api.DELETE("/products/:id", middlewares.AuthenticationMiddleware, controllers.DeleteProduct)
	}
}
