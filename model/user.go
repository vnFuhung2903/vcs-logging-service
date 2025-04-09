package model

import (
	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Password  string
	Email     string         `gorm:"unique;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Logs      []Log          `gorm:"foreignKey:UserId;"`
}
