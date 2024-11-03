package main

type Item struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Quantity int     `json:"quantity"`
	Rating   float64 `json:"rating"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
