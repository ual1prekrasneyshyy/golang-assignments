package main

import "fmt"

func main() {
	var dayOfWeek int

	fmt.Println("Enter number of day of week:")

	fmt.Scan(&dayOfWeek)

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("You had to enter a number from 1 to 7")
	}

	//switch {
	//case dayOfWeek == 1:
	//	fmt.Println("Monday")
	//case dayOfWeek == 2:
	//	fmt.Println("Tuesday")
	//case dayOfWeek == 3:
	//	fmt.Println("Wednesday")
	//case dayOfWeek == 4:
	//	fmt.Println("Thursday")
	//case dayOfWeek == 5:
	//	fmt.Println("Friday")
	//case dayOfWeek == 6:
	//	fmt.Println("Saturday")
	//case dayOfWeek == 7:
	//	fmt.Println("Sunday")
	//default:
	//	fmt.Println("You had to enter a number from 1 to 7")
	//}
}
