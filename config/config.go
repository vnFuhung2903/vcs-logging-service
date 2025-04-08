package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vnFuhung2903/vcs-logging-service/model"
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

	fmt.Println("Connected to postgresql db")
	err = db.AutoMigrate(&model.User{}, &model.Log{})
	if err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
	return db
}

func ConnectServices(db *gorm.DB) service.UserService {
	userRepo := repository.NewUserRepository(db)
	return service.NewUserService(&userRepo)
}
