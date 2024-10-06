package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Age  int
}

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected to database")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserById(id string) (User, error) {
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

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users(id, name, age) VALUES($1, $2, $3)", user.Id, user.Name, user.Age)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("User has successfully been created")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)

	userId := httpParams["id"]

	_, err := GetUserById(userId)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	var newUserData User

	err = json.NewDecoder(r.Body).Decode(&newUserData)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", newUserData.Name, newUserData.Age, userId)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(newUserData)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	userId := httpParams["id"]

	_, err := GetUserById(userId)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM users WHERE id=$1", userId)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("User Deleted")
	}
}

func main() {
	Connect()

	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Server starting on port 8000...")
	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal(err)
	}
}
