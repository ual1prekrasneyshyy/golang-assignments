package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	var err error

	db, err = gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"),
		&gorm.Config{},
	)

	// If error occcures during connection to database, the following message will be logged and the program will be stopped
	if err != nil {
		log.Fatal("An error occured during connection to database")
		os.Exit(1)
	} else {
		log.Println("Successfully connected to database")
	}
}

// "Getter". It will return the instance of connection to database
func GetConnection() *gorm.DB {
	return db
}

func MakeMigration() {
	GetConnection().AutoMigrate(&Role{})
	GetConnection().AutoMigrate(&User{})
	GetConnection().AutoMigrate(&Post{})

	log.Println("Tables has been created")
}
