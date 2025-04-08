package main

import (
	"log"
	"os"
	"time"

	"github.com/vnFuhung2903/vcs-logging-service/config"
	"gorm.io/gorm"
)

func addTrigger(db *gorm.DB) {
	sqlBytes, err := os.ReadFile("migration/logs.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	execTrigger := db.Exec(string(sqlBytes))
	if execTrigger.Error != nil {
		log.Fatalf("Failed to execute trigger SQL: %v", execTrigger.Error)
	}
}

func main() {
	db := config.ConnectPostgresDb()

	// addTrigger(db)

	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		startTime := time.Now().Unix()
		for i := range 500 {
			email := string(rune(i)) + "@gmail.com"
			_, err := userService.Register(email, string(rune(i)))
			if err != nil {
				return err
			}
		}
		log.Printf("Update 500 records in %v", time.Now().Unix()-startTime)
		return nil
	})
	if err != nil {
		log.Println("Transaction error: ", err)
	}
}
