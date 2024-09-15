package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello! My name is %s. I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{"Uali", 21}
	person.Greet()
}
