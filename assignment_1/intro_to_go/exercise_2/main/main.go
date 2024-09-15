package main

import "fmt"

func main() {
	var a int = 5
	fmt.Printf("a is %T.", a)
	fmt.Println()
	fmt.Printf("a = %d.", a)
	fmt.Println()

	b := 4
	fmt.Printf("b is %T, b = %d.", b, b)
	fmt.Println()

	var c float64 = 4.32
	fmt.Printf("c is %T, c = %f. \n", c, c)

	d := 0.43
	fmt.Printf("d is %T, d = %f. \n", d, d)

	var e complex64 = 6 + 7i
	fmt.Printf("e is %T, e = %g. \n", e, e)

	f := 5 - 6i
	fmt.Printf("f is %T, f = %g. \n", f, f)

	var g string = "Hello, world"
	fmt.Printf("g is %T, g = %s. \n", g, g)

	h := "subject"
	fmt.Printf("h is %T, h = %s. \n", h, h)

	var i bool = true
	fmt.Printf("i is %T, i = %t. \n", i, i)

	j := false
	fmt.Printf("j is %T, j = %v. \n", j, j)

}
