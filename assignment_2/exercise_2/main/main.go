package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Id   int
	Name string
	Age  int
}

func main() {
	// First of all, I make connection to database server via GORM
	db, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=task2 sslmode=disable"),
		&gorm.Config{},
	)

	// If some errors occur, they will be logged to the console
	if err != nil {
		log.Fatal(err)
	}

	// I make migration
	err = db.AutoMigrate(&User{})

	// If some errors occur, they will be logged to the console
	if err != nil {
		log.Fatal(err)
	}

	// This is a menu, that continuously will ask from user what to do - to insert new user or to print users list
	for true {
		fmt.Println("PRESS [1] TO INSERT USER")
		fmt.Println("PRESS [2] TO VIEW ALL USERS")
		fmt.Println("PRESS [3] TO QUIT")

		var choice int

		fmt.Scan(&choice)

		// If user choose 1, the program will ask to insert data of user and insert them to the database
		if choice == 1 {
			var name string
			var age int
			// User inserts name & age
			fmt.Print("Insert name: ")
			fmt.Scan(&name)
			fmt.Print("Insert age: ")
			fmt.Scan(&age)

			// The INSERT query will be executed to the database
			db.Create(&User{Name: name, Age: age})

		} else if choice == 2 {
			// If user choose 2, the program will query list of all users from database. After this list will be printed in the console.

			var users []User
			// Query list of all users
			db.Find(&users)

			// Printing all user data via loop
			for i := 0; i < len(users); i++ {
				fmt.Printf("ID: %d, Name: %s, Age: %d \n", users[i].Id, users[i].Name, users[i].Age)
			}
		} else if choice == 3 {
			break
		} else {
			break
		}
	}

}
