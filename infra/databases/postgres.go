package databases

import (
	"fmt"
	"log"

	"github.com/vnFuhung2903/vcs-logging-services/models"
	"github.com/vnFuhung2903/vcs-logging-services/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDb(env env.Env) *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", env.PostgresUser, env.PostgresPassword, env.PostgresName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}
	fmt.Println("Connected to postgresql db")

	err = db.AutoMigrate(&models.User{}, &models.Log{})
	if err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
	return db
}
