package model

import (
	"time"
)

type Log struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null;index"`
	Operation string `gorm:"not null;index"`
	Collumn   string
	OldData   string
	NewData   string
	CreatedAt time.Time
}
