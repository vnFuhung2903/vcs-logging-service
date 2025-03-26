package model

type Wallet struct {
	Id           uint `gorm:"primaryKey"`
	UserId       uint `gorm:"not null"`
	WalletNumber uint `gorm:"unique"`
	Balance      uint `gorm:"default:0"`
}
