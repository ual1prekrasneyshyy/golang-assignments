package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Id   int
	Name string
	Age  int
}

var db *gorm.DB

func Connect() {
	var err error

	db, err = gorm.Open(
		postgres.Open("host = localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	db.AutoMigrate(&User{})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&user)

	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)

	userId := httpParams["id"]

	var user User

	db.First(&user, userId)

	if user.Id > 0 {
		json.NewDecoder(r.Body).Decode(&user)

		db.Updates(&user)

		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode("User not found")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)

	userId := httpParams["id"]

	var user User

	db.First(&user, userId)

	if user.Id > 0 {
		db.Delete(&user)

		json.NewEncoder(w).Encode("User deleted")
	} else {
		json.NewEncoder(w).Encode("User not found")
	}
}

func main() {
	Connect()
	Migrate()

	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/user", AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Server starting on port 8001...")
	err := http.ListenAndServe(":8001", router)

	if err != nil {
		log.Fatal(err)
	}
}
