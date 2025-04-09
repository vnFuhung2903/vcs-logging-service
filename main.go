package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/vnFuhung2903/vcs-logging-service/config"
	"github.com/vnFuhung2903/vcs-logging-service/model"
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

func checkAdddUsers(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		startTime := time.Now().UnixMilli()
		for i := range 500 {
			email := fmt.Sprint(i, "@gmail.com")
			_, err := userService.Register(email, email)
			if err != nil {
				return err
			}
		}
		log.Printf("Insert 500 records in %fs", float64(time.Now().UnixMilli()-startTime)/1000)
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

		startTime := time.Now().UnixMilli()
		for i := range 500 {
			newPassword := fmt.Sprint(i)
			err := userService.Update(user, "password", newPassword)
			if err != nil {
				return err
			}
		}
		log.Printf("Update 500 records in %fs", float64(time.Now().UnixMilli()-startTime)/1000)
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	}
}

func checkDeleteUser(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		userService := config.ConnectServices(tx)
		startTime := time.Now().UnixMilli()
		for i := range 500 {
			email := fmt.Sprint(i, "@gmail.com")
			user, err := userService.FindByEmail(email)
			if err != nil {
				return err
			}

			err = userService.Delete(user)
			if err != nil {
				return err
			}
		}
		log.Printf("Delete 500 records in %fs", float64(time.Now().UnixMilli()-startTime)/1000)
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction error: %v", err)
	}
}

func writeLogsToES(db *gorm.DB, es *elasticsearch.Client, lastTime time.Time) {
	var rows []model.Log
	res := db.Find(&rows).Where("create_at > ?", lastTime)
	if res.Error != nil {
		log.Fatalf("Reading logs error: %s", res.Error)
	}

	log.Println("Number of logs: ", len(rows))
	startTime := time.Now().UnixMilli()
	for _, row := range rows {
		data, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}

		res, err := es.Index(
			fmt.Sprint(lastTime),
			bytes.NewReader(data),
			es.Index.WithContext(context.Background()),
		)
		if err != nil {
			log.Fatalf("ES Indexing error: %v", err)
			break
		}
		if res.IsError() {
			log.Fatalf("ES Indexing error: %v", res.String())
			break
		}
		defer res.Body.Close()
	}
	log.Printf("Delete 500 records in %fs", float64(time.Now().UnixMilli()-startTime)/1000)
}

func main() {
	lastTime := time.Now()
	db := config.ConnectPostgresDb()
	addTrigger(db)
	checkAdddUsers(db)
	checkUpdateUser(db)
	checkDeleteUser(db)

	es := config.ConnectESDb()
	writeLogsToES(db, es, lastTime)
}
