package main

import (
	"ecommerce-go-be/database/seeders"
	"ecommerce-go-be/internal/config"
	"ecommerce-go-be/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db := config.ConnectDB()

	// Migrate the database
	config.MigrateDB(db)

	// Seed the database
	seeders.SeedAll(db)

	// Initialize the router
	router := gin.Default()

	// Setup the routes
	routes.SetupRouter(router)

	// Run the server
	router.Run(":8080")
}
