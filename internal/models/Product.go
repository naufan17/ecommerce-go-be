package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"type:primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Price       float64   `json:"price" gorm:"type:decimal;not null"`
	Quantity    int       `json:"quantity" gorm:"type:integer;not null;default:1"`
	CategoryID  uint      `json:"category_id" gorm:"type:integer;not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`

	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
