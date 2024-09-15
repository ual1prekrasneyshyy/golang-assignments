package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	fmt.Println(sum)
}
