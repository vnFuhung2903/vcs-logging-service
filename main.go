package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/vnFuhung2903/vcs-logging-service/config"
	"github.com/vnFuhung2903/vcs-logging-service/model"
	"gorm.io/gorm"
)

func main() {
	db := config.ConnectPostgresDb()
	fmt.Println("Connected to postgresql db")
	err := db.AutoMigrate(&model.User{}, &model.Wallet{})
	if err != nil {
		log.Println("Migrate error: ", err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		authService, userService, walletService := config.ConnectServices(tx)
		emailTest := "vcs123@test"
		passwordTest := "123456789Aa@"
		nameTest := "VCSTest"
		userCheck, err := userService.GetUserByEmail(emailTest)
		if err != nil {
			return err
		}

		var user *model.User
		if len(userCheck) == 0 {
			user, err = userService.Register(emailTest, passwordTest, nameTest)
		} else {
			user, err = authService.Login(emailTest, passwordTest)
		}
		if err != nil {
			return err
		}

		wallet, err := walletService.CreateNewWallet(user.Id, rand.Uint())
		if wallet != nil {
			log.Println("Wallet created: ", wallet.WalletNumber)
		}
		return err
	})
	if err != nil {
		log.Println("Transaction error: ", err)
	}
}
