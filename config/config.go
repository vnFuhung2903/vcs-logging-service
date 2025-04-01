package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
	"github.com/vnFuhung2903/vcs-logging-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func ConnectServices(db *gorm.DB) (service.AuthService, service.UserService, service.WalletService) {
	userRepo := repository.NewUserRepository(db)
	walletRepo := repository.NewWalletRepository(db)

	return service.NewAuthService(&userRepo), service.NewUserService(&userRepo), service.NewWalletService(&walletRepo)
}
