package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/vnFuhung2903/vcs-logging-services/infra/databases"
	"github.com/vnFuhung2903/vcs-logging-services/infra/messages"
	"github.com/vnFuhung2903/vcs-logging-services/pkg/env"
	"github.com/vnFuhung2903/vcs-logging-services/usecases/repositories"
	"github.com/vnFuhung2903/vcs-logging-services/usecases/services"
	"gorm.io/gorm"
)

func connectUserService(db *gorm.DB) services.UserService {
	userRepo := repositories.NewUserRepository(db)
	return services.NewUserService(userRepo)
}

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

func checkAddUsers(db *gorm.DB, workers uint) {
	var wg sync.WaitGroup
	emails := make(chan string, 500)

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for email := range emails {
				err := db.Transaction(func(tx *gorm.DB) error {
					userService := connectUserService(tx)
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
}

func checkUpdateUser(db *gorm.DB, workers uint) {
	var wg sync.WaitGroup
	passwords := make(chan string, 500)

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for password := range passwords {
				err := db.Transaction(func(tx *gorm.DB) error {
					userService := connectUserService(tx)
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
}

func checkDeleteUser(db *gorm.DB, workers uint) {
	var wg sync.WaitGroup
	emails := make(chan string, 500)

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for email := range emails {
				err := db.Transaction(func(tx *gorm.DB) error {
					userService := connectUserService(tx)
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
}

func main() {
	env, err := env.LoadConfig("./")
	if err != nil {
		log.Fatal("Cannot load env variables")
	}
	db := databases.ConnectPostgresDb(env)
	kafkaWriter := messages.ConnectKafkaWriter(fmt.Sprintf("%s:9092", env.KafkaBrokerAddress), "logstash")
	defer kafkaWriter.Close()
	logRepo := repositories.NewLogRepository(db)
	logService := services.NewLogService(logRepo, kafkaWriter, 5, 500)

	addTrigger(db)
	checkAddUsers(db, 5)
	// checkUpdateUser(db, 5)
	// checkDeleteUser(db, 5)

	startTime := time.Now()
	err = logService.Process()
	log.Printf("Process 500 records in %v", time.Since(startTime))
	if err != nil {
		log.Print(err)
	}

	cron := cron.New(cron.WithSeconds())
	_, err = cron.AddFunc("* */10 * * * *", func() {
		err = logService.DeleteProcessedLogs()
		if err != nil {
			log.Print(err)
		}
	})

	cron.Start()
	select {}
}
