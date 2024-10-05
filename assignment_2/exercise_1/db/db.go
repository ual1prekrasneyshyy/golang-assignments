package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=postgres password=postgres")

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected to database")
}
