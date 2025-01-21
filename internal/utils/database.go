package utils

import (
	"ecommerce-go-be/internal/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.ConnectDB()
}
