package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Password  string
	Email     string         `gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Logs      []Log          `gorm:"foreignKey:UserId;"`
}
