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
		"host=localhost port=5432 user=postgres dbname=task1advanced password=postgres sslmode=disable")
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
	_, err := db.Exec("CREATE TABLE users(id SERIAL, name VARCHAR(30) UNIQUE NOT NULL, age INT NOT NULL, PRIMARY KEY(id));")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table users has successfully been create")
}

func InsertUsers(users []User) bool {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	// Defer rollback in case of error
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to: %v", r)
		}
	}()
	// Insert each user
	for _, user := range users {
		_, err = tx.Exec("INSERT INTO users(name, age) VALUES(?, ?)", user.Name, user.Age)

		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return false
		}
	}
	// Commit transaction if all insertions were successful
	tx.Commit()
	// return true if all users successfully has been added
	return true
}

func LoadAllUsersByMinAge(age int) []User {
	var users []User

	rows, err := db.Query("SELECT id, name, age FROM users WHERE age >= $1;", age)

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

func LoadAllUsersAndPaging(page, size int) []User {
	var users []User

	rows, err := db.Query("SELECT id, name, age FROM users OFFSET $1 LIMIT $2;", (page-1)*size, size)

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

func GetUserById(id int) (User, error) {
	var user User

	row, err := db.Query("SELECT * FROM users WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)

		return user, err
	}

	if row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Age)

		if err != nil {
			log.Fatal(err)

			return User{}, err
		}

		return user, nil
	}

	return user, fmt.Errorf("User Not Found!")
}

func UpdateUser(id int, user User) bool {

	_, err := GetUserById(id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = db.Exec("UPDATE users SET name=?, age=? WHERE id=?", user.Name, user.Age, id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func DeleteUser(id int) bool {

	_, err := GetUserById(id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func main() {
	connect()
	//// Then I call function to create table (only once)
	CreateTable()

	// This is a menu, that continuously will ask from user what to do - to insert new user or to print users list
	for true {
		fmt.Println("PRESS [1] TO INSERT TWO USERS")
		fmt.Println("PRESS [2] TO VIEW ALL USERS")
		fmt.Println("PRESS [3] TO UPDATE USER")
		fmt.Println("PRESS [4] TO DELETE USER")
		fmt.Println("PRESS [0] TO QUIT")

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

			u1 := &User{Name: name, Age: age}

			// User inserts name & age
			fmt.Print("Insert name: ")
			fmt.Scan(&name)
			fmt.Print("Insert age: ")
			fmt.Scan(&age)

			u2 := &User{Name: name, Age: age}

			InsertUsers([]User{*u1, *u2})

		} else if choice == 2 {
			// If user choose 2, the program will query list of all users from database. After this list will be printed in the console.

			users := LoadAllUsersByMinAge(21)
			// Query list of all users

			// Printing all user data via loop
			for i := 0; i < len(users); i++ {
				fmt.Printf("ID: %d, Name: %s, Age: %d \n", users[i].Id, users[i].Name, users[i].Age)
			}
		} else if choice == 3 {
			var id int

			fmt.Print("Insert id: ")

			fmt.Scan(&id)
			var name string
			var age int
			// User inserts name & age
			fmt.Print("Insert name: ")
			fmt.Scan(&name)
			fmt.Print("Insert age: ")
			fmt.Scan(&age)

			UpdateUser(id, User{Name: name, Age: age})

		} else if choice == 4 {
			var id int

			fmt.Print("Insert id: ")

			fmt.Scan(&id)

			DeleteUser(id)
		} else if choice == 0 {
			break
		} else {
			break
		}
	}
}
