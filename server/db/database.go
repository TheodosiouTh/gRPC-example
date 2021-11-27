package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init(host, dbname, user, password string, port int) error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err = connectWithRetries(3, connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection to DB established")
	migrateModels()
	return nil
}

func connectWithRetries(retries int, connectionString string) (*gorm.DB, error) {
	connection, er := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	retries--
	for er != nil || retries == 0 {
		time.Sleep(5 * time.Second)
		connection, er = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		retries--
	}

	if er != nil {
		return nil, er
	}
	return connection, nil
}

func migrateModels() {
	db.AutoMigrate(&Task{})
}
