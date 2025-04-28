package model

import (
	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey"`
	Email     string         `gorm:"unique;index;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Logs      []Log          `gorm:"foreignKey:UserId;"`
	Password  string
}
