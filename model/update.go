package model

import "time"

type Update struct {
	Id        uint      `gorm:"primaryKey"`
	Table     string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
	DataId    uint      `gorm:"not null"`
}
