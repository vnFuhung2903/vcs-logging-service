package model

type User struct {
	Id       uint `gorm:"primaryKey"`
	Password string
	Email    string `gorm:"unique;not null"`
}
