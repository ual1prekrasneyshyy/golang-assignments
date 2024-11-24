package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	Connect()
	MakeMigration()

	router := mux.NewRouter()

	router.HandleFunc("/posts", GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", CheckJWT(AddPost)).Methods("POST")
	router.HandleFunc("/posts/{id}", GetPostById).Methods("GET")
	router.HandleFunc("/posts/{id}", CheckJWT(CheckAdminRights(UpdateUser))).Methods("PUT", "PATCH")
	router.HandleFunc("/posts/{id}", CheckJWT(CheckAdminRights(DeleteUser))).Methods("DELETE")

	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/sign-up", SignUp).Methods("POST")

	fmt.Println("Server starting on port 8000...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {

	}
}
