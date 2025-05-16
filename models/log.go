package models

import (
	"time"
)

type Log struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null;index"`
	Operation string    `gorm:"not null;index"`
	Processed bool      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Collumn   string
	OldData   string
	NewData   string
}
