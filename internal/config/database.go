package config

import (
	"log"
	"os"

	"ecommerce-go-be/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Connected to database")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	err = db.AutoMigrate(&models.Category{})
	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	} else {
		log.Println("Migrated database")
	}
}
