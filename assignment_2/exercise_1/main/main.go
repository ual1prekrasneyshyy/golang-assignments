package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// This will be global variable for other functions
var db *sql.DB

// Structure User
type User struct {
	Id   int
	Name string
	Age  int
}

// connection to database
func connect() {
	var err error
	db, err = sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	// I hade some problems with SSL mode, so I have disabled it for the test.

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected to database")
}

// CreateTable This function creates table
func CreateTable() {
	_, err := db.Exec("CREATE TABLE users(id INT, name VARCHAR(30), age INT);")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table users has successfully been create")
}

func InsertUser(u User) bool {
	_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3);", u.Id, u.Name, u.Age)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func LoadAllUsers() []User {
	var users []User

	rows, err := db.Query("SELECT id, name, age FROM users;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User

		err = rows.Scan(&u.Id, &u.Name, &u.Age)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

	return users

}

func main() {
	// Firstly, I connect to database
	connect()
	//// Then I call function to create table (only once)
	//CreateTable()
	//
	//id1 := InsertUser(User{1, "Uali", 21})
	//id2 := InsertUser(User{2, "Alex", 20})
	//
	//if id1 {
	//	fmt.Println("First user has successfully been added")
	//}
	//
	//if id2 {
	//	fmt.Println("Second user has successfully been added")
	//}

	users := LoadAllUsers()

	for i := 0; i < len(users); i++ {
		fmt.Printf("ID: %d, Name: %s, Age: %d \n", users[i].Id, users[i].Name, users[i].Age)
	}
}
