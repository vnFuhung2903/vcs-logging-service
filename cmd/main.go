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

func checkAddUsers(db *gorm.DB) {
	var wg sync.WaitGroup
	emails := make(chan string, 500)
	startTime := time.Now()

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for email := range emails {
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
		}()
	}
	for i := range 500 {
		email := fmt.Sprint(i, "@gmail.com")
		emails <- email
	}
	close(emails)
	wg.Wait()
	log.Printf("Insert 500 records in %v", time.Since(startTime))
}

func checkUpdateUser(db *gorm.DB) {
	var wg sync.WaitGroup
	passwords := make(chan string, 500)
	startTime := time.Now()

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for password := range passwords {
				err := db.Transaction(func(tx *gorm.DB) error {
					userService := config.ConnectServices(tx)
					user, err := userService.FindByEmail("1@gmail.com")
					if err != nil {
						return err
					}

					err = userService.Update(user, "password", password)
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					log.Fatalf("Transaction error: %v", err)
				}
			}
		}()
	}
	for range 500 {
		passwords <- "password"
	}
	close(passwords)
	wg.Wait()
	log.Printf("Update 500 records in %v", time.Since(startTime))
}

func checkDeleteUser(db *gorm.DB) {
	var wg sync.WaitGroup
	emails := make(chan string, 500)
	startTime := time.Now()

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for email := range emails {
				err := db.Transaction(func(tx *gorm.DB) error {
					userService := config.ConnectServices(tx)
					err := userService.Delete(email)
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					log.Fatalf("Transaction error: %v", err)
				}
			}
		}()
	}
	for i := range 500 {
		email := fmt.Sprint(i, "@gmail.com")
		emails <- email
	}
	close(emails)
	wg.Wait()
	log.Printf("Delete 500 records in %v", time.Since(startTime))
}

func main() {
	db := config.ConnectPostgresDb()
	addTrigger(db)
	checkAddUsers(db)
	checkUpdateUser(db)
	checkDeleteUser(db)
	deleteLogs(db)
}
