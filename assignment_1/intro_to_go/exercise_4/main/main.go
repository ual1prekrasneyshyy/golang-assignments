package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func divide(numerator, denominator int) (int, int) {
	quotient := numerator / denominator
	reminder := numerator % denominator
	return quotient, reminder
}

func main() {
	sum := add(45, 65)
	fmt.Printf("Sum is %d. \n", sum)

	s1 := "Hello"
	s2 := "World"
	text1, text2 := swap(s1, s2)

	fmt.Println(text1 + " " + text2)

	q, r := divide(30, 7)
	fmt.Printf("30 / 7: quotient=%d, reminder=%d. \n", q, r)

}
