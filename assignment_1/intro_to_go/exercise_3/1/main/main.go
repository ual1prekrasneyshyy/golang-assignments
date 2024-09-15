package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Println("Number is positive")
	} else if a < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}
}
