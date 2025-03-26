package model

type Transaction struct {
	Id         uint `gorm:"primaryKey"`
	From       uint `gorm:"not null;index:,using:gin"`
	To         uint `gorm:"not null"`
	Amount     uint `gorm:"default:0"`
	CreateTime uint `gorm:"index:,using:brin"`
}
