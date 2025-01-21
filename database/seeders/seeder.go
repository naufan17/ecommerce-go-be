package seeders

import (
	"ecommerce-go-be/internal/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seederUsers(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	users := []models.User{
		{Name: "John Doe", Email: "jhon@example.com", Password: string(hashedPassword)},
		{Name: "Jane Doe", Email: "jane@example.com", Password: string(hashedPassword)},
		{Name: "Mark Doe", Email: "mark@example.com", Password: string(hashedPassword)},
		{Name: "Alice Doe", Email: "alice@example.com", Password: string(hashedPassword)},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
}

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Laptop"},
		{Name: "Notebook"},
		{Name: "Tablet"},
		{Name: "Personal Computer"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Fatalf("Failed to seed categories: %v", err)
	}
}

func seedProducts(db *gorm.DB) {
	products := []models.Product{
		{Name: "Lenovo Ideapad", Description: "Intel Core i5, 8GB RAM, 256GB SSD, 14 inch", Price: 8000000.00, Quantity: 20, CategoryID: 1, UserID: 1},
		{Name: "Lenovo Yoga", Description: "Intel Core i7, 16GB RAM, 1TB SSD, 14.5 inch", Price: 20000000.00, Quantity: 15, CategoryID: 1, UserID: 2},
		{Name: "Asus VivoBook", Description: "Intel Core i7, 16GB RAM, 512GB SSD, 14 inch", Price: 15000000.00, Quantity: 20, CategoryID: 1, UserID: 3},
		{Name: "Asus Zenbook", Description: "Intel Core Ultra 7, 32GB RAM, 1TB SSD, 14 inch", Price: 17500000.00, Quantity: 15, CategoryID: 1, UserID: 4},
		{Name: "Acer Aspire", Description: "Intel Core i5, 8GB RAM, 512GB SSD, 14 inch", Price: 10000000.00, Quantity: 20, CategoryID: 1, UserID: 2},
		{Name: "Acer Swift", Description: "Intel Core Ultra 5, 16GB RAM, 512GB SSD, 14 inch", Price: 12500000.00, Quantity: 15, CategoryID: 1, UserID: 4},
		{Name: "Macbook M4 Pro", Description: "M4, 16GB RAM, 512GB SSD, 14 inch", Price: 28000000.00, Quantity: 10, CategoryID: 1, UserID: 1},
		{Name: "Macbook M3 Air", Description: "M3, 16GB RAM, 256GB SSD, 13 inch", Price: 16500000.00, Quantity: 10, CategoryID: 1, UserID: 2},
		{Name: "HP Pavilion", Description: "Intel Core i5, 16GB RAM, 512GB SSD, 14 inch", Price: 10000000.00, Quantity: 15, CategoryID: 1, UserID: 3},
		{Name: "Microsoft Surface", Description: "Intel Core i5, 16GB RAM, 512GB SSD, 13.3 inch", Price: 15000000.00, Quantity: 5, CategoryID: 1, UserID: 1},
		{Name: "Asus ROG Zephyrus", Description: "Intel Core Ultra 9, 32GB RAM, 1TB SSD, 15.6 inch", Price: 35000000.00, Quantity: 5, CategoryID: 1, UserID: 3},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatalf("Failed to seed products: %v", err)
	}
}

func SeedAll(db *gorm.DB) {
	seederUsers(db)
	seedCategories(db)
	seedProducts(db)
}
