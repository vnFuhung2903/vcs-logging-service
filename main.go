package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/vnFuhung2903/vcs-logging-service/config"
	"gorm.io/gorm"
)

func addTrigger(db *gorm.DB) {
	sqlBytes, err := os.ReadFile("migration/add_trigger.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	execTrigger := db.Exec(string(sqlBytes))
	if execTrigger.Error != nil {
		log.Fatalf("Failed to execute trigger SQL: %v", execTrigger.Error)
	}
}

func deleteLogs(db *gorm.DB) {
	sqlBytes, err := os.ReadFile("migration/del_logs.sql")
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	execTrigger := db.Exec(string(sqlBytes))
	if execTrigger.Error != nil {
		log.Fatalf("Failed to execute trigger SQL: %v", execTrigger.Error)
	}
}

func checkAddUsers(db *gorm.DB, email string) {
	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		_, err := userService.Register(email, email)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	}
}

func checkUpdateUser(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		user, err := userService.FindByEmail("1@gmail.com")
		if err != nil {
			return err
		}

		startTime := time.Now()
		for i := range 500 {
			newPassword := fmt.Sprint(i)
			err := userService.Update(user, "password", newPassword)
			if err != nil {
				return err
			}
		}
		log.Printf("Update 500 records in %v", time.Since(startTime))
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	}
}

func checkDeleteUser(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		startTime := time.Now()
		for i := range 500 {
			email := fmt.Sprint(i, "@gmail.com")
			err := userService.Delete(email)
			if err != nil {
				return err
			}
		}
		log.Printf("Delete 500 records in %v", time.Since(startTime))
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	}
}

func main() {
	db := config.ConnectPostgresDb()
	addTrigger(db)
	deleteLogs(db)
	var wg sync.WaitGroup
	emails := make(chan string, 500)

	startTime := time.Now()
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for email := range emails {
				checkAddUsers(db, email)
			}
		}()
	}
	for i := range 500 {
		email := fmt.Sprint(i, "@gmail.com")
		emails <- email
	}
	close(emails)
	wg.Wait()
	log.Printf("Insert 500 records in %v", time.Since(startTime))
	// checkUpdateUser(db)
	// checkDeleteUser(db)
}
