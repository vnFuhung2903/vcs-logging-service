package model

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Password string
	Email    string   `gorm:"unique;not null"`
	Wallets  []Wallet `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`
}
