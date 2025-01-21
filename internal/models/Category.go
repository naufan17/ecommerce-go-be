package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryID uint      `json:"category_id" gorm:"primaryKey;autoIncrement;not null"`
	Name       string    `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
