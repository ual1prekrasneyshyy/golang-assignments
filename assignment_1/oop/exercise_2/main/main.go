package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Department string
}

func (e Employee) Work() {
	fmt.Printf("Manager with name = %s and ID = %s. \n", e.Name, e.ID)
}

func main() {
	employee := Employee{"Uali", "21B031234"}
	employee.Work()
}
