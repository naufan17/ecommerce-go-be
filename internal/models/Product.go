package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID   uint      `json:"product_id" gorm:"type:primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Price       float64   `json:"price" gorm:"type:decimal;not null"`
	Quantity    int       `json:"quantity" gorm:"type:integer;not null;default:1"`
	CategoryID  uint      `json:"category_id" gorm:"type:integer;not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Category Category `gorm:"foreignKey:CategoryID;references:CategoryID"`

	User User `gorm:"foreignKey:UserID;references:ProductID"`
}
