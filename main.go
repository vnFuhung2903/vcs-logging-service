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
		startTime := time.Now()
		for i := range 500 {
			email := fmt.Sprint(i, "@gmail.com")
			_, err := userService.Register(email, email)
			if err != nil {
				return err
			}
		}
		log.Printf("Insert 500 records in %v", time.Since(startTime))
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
			user, err := userService.FindByEmail(email)
			if err != nil {
				return err
			}

			err = userService.Delete(user)
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

func checkWriteLogsToES(db *gorm.DB, es *elasticsearch.Client, lastTime string) {
	var rows []model.Log
	res := db.Find(&rows).Where("create_at > ?", lastTime)
	if res.Error != nil {
		log.Fatalf("Reading logs error: %s", res.Error)
	}

	startTime := time.Now()
	var buf bytes.Buffer
	for _, row := range rows {
		meta := fmt.Appendf(nil, `{ "index" : { "_index" : "%s" } }\n`, lastTime)
		data, err := json.Marshal(row)
		if err != nil {
			log.Fatalf("Json marshaling error: %v", err)
		}
		data = append(data, byte('\n'))

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}

	bulkRes, err := es.Bulk(
		bytes.NewReader(buf.Bytes()),
		es.Bulk.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Bulk indexing error: %v", err)
	}
	if bulkRes.IsError() {
		log.Fatalf("Bulk indexing response error: %s", bulkRes.String())
	}
	defer bulkRes.Body.Close()
	log.Printf("Write 500 records to ES in %v", time.Since(startTime))
}

func main() {
	lastTime := time.Now().Format("2004-03-29")
	db := config.ConnectPostgresDb()
	addTrigger(db)
	checkAdddUsers(db)
	checkUpdateUser(db)
	checkDeleteUser(db)

	es := config.ConnectESDb()
	checkWriteLogsToES(db, es, lastTime)
}
