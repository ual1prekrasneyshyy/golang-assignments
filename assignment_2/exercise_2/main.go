package exercise_2

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int
	Name string
	Age  int
}

func connect() {
	var err error
	db, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=task2 sslmode=disable"),
		&gorm.Config{},
	)
}
